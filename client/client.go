package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/levigross/grequests"
	"github.com/maczh/chinaums-invoice-sdk/util"
	"github.com/maczh/mgo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Client 银联商务发票服务客户端
type Client struct {
	mgoClient *mgo.Session // MongoDB会话
	dbName    string       // 数据库名称
	config    *Config
}

var logger = logrus.New()

var client *Client

// Config 客户端配置
type Config struct {
	BaseURL        string // 接口基础地址（测试/生产）
	MsgSrc         string // 消息来源
	SignKey        string // 签名密钥
	MerchantID     string // 银商商户号
	TerminalID     string // 银商终端号
	Timeout        time.Duration
	MongoURL       string // MongoDB连接URL
	CollectionName string // 接口调用日志集合
	IsProd         bool   // 是否生产环境
}

// NewClient 创建客户端
func NewClient(env string, cnf ...Config) (*Client, error) {
	var err error
	if len(cnf) > 0 && client == nil {
		client = &Client{
			config: &cnf[0],
		}
		if cnf[0].MongoURL != "" {
			client.mgoClient, err = mgo.Dial(cnf[0].MongoURL)
			if err != nil {
				return client, err
			}
			client.dbName = extractDBName(cnf[0].MongoURL)
		}
		return client, nil
	}
	if env == "" {
		env = "prod"
	}
	config, ok := configs[env]
	if !ok {
		return nil, errors.New("invalid environment")
	}
	if config.BaseURL == "" {
		return nil, errors.New("baseURL is required")
	}
	if config.MsgSrc == "" || config.SignKey == "" {
		return nil, errors.New("msgSrc and signKey are required")
	}

	if client != nil {
		return client, nil
	}

	client = &Client{
		config: &config,
	}
	if config.MongoURL != "" {
		client.mgoClient, err = mgo.Dial(config.MongoURL)
		if err != nil {
			return client, err
		}
		client.dbName = extractDBName(config.MongoURL)
	}
	logger.SetOutput(os.Stdout)        // 设置日志输出到标准输出
	logger.SetLevel(logrus.DebugLevel) // 设置日志级别为Debug

	return client, nil
}

// SendRequest 发送请求（通用封装）
func SendRequest[R any, T any](req R, resp *T) error {
	if client == nil {
		client, _ = NewClient("prod")
	}
	// 转换为map用于签名
	reqMap, err := util.StructToMap(req)
	if err != nil {
		return err
	}
	// 填充公共请求参数
	reqMap["msgId"] = uuid.NewString()
	reqMap["msgSrc"] = client.config.MsgSrc
	reqMap["requestTimestamp"] = time.Now().Format("2006-01-02 15:04:05")
	if reqMap["merchantId"] == nil || reqMap["merchantId"] == "" {
		reqMap["merchantId"] = client.config.MerchantID
	}
	if reqMap["terminalId"] == nil || reqMap["terminalId"] == "" {
		reqMap["terminalId"] = client.config.TerminalID
	}

	// 生成签名
	reqMap["sign"] = util.Sign(reqMap, client.config.SignKey)

	// 记录请求日志
	postLog := Postlog{
		Id:      primitive.NewObjectID(),
		ReqTime: time.Now().Format("2006-01-02 15:04:05.000"),
		ReqBody: reqMap,
	}

	// 输出请求日志
	// logger.Debugf("Request: %s", util.ToJsonString(reqMap))

	// 发送请求
	res, err := grequests.DoRegularRequest(http.MethodPost, client.config.BaseURL, &grequests.RequestOptions{
		JSON: reqMap,
	})
	postLog.RespTime = time.Now().Format("2006-01-02 15:04:05.000") // 响应时间
	if err != nil {
		// 记录响应日志
		postLog.RespBody = err.Error()
		if client.mgoClient != nil {
			_ = client.mgoClient.DB(client.dbName).C(client.config.CollectionName).Insert(postLog)
		}
		return err
	}

	// 输出响应日志
	// logger.Debugf("Response: %s", res.String())

	// 校验响应签名
	// respMap, _ := util.StructToMap(resp)
	// sign := respMap["sign"].(string)
	// delete(respMap, "sign")
	// if util.Sign(respMap, client.config.SignKey) != sign {
	// 	return errors.New("response sign verify failed")
	// }
	err = json.Unmarshal(res.Bytes(), resp)
	if err != nil {
		postLog.RespBody = map[string]interface{}{
			"response": res.String(),
			"error":    err.Error(),
		}
		if client.mgoClient != nil {
			_ = client.mgoClient.DB(client.dbName).C(client.config.CollectionName).Insert(postLog)
		}
		return err
	}
	var respMap map[string]interface{}
	err = json.Unmarshal(res.Bytes(), &respMap)
	if err != nil {
		postLog.RespBody = map[string]interface{}{
			"response": res.String(),
			"error":    err.Error(),
		}
		if client.mgoClient != nil {
			_ = client.mgoClient.DB(client.dbName).C(client.config.CollectionName).Insert(postLog)
		}
		return err
	}
	postLog.RespBody = resp
	if client.mgoClient != nil {
		_ = client.mgoClient.DB(client.dbName).C(client.config.CollectionName).Insert(postLog)
	}
	// 检查响应状态
	if respMap["resultCode"].(string) != "SUCCESS" {
		return errors.New(respMap["resultMsg"].(string))
	}

	return nil
}

// 从MongoDB连接字符串中提取数据库名称
func extractDBName(connectionString string) string {
	str := strings.Split(connectionString, "?")[0]
	ss := strings.Split(str, "/")
	return ss[len(ss)-1]
}

type Postlog struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	ReqTime  string             `json:"reqTime" bson:"reqTime"`   // 请求时间
	RespTime string             `json:"respTime" bson:"respTime"` // 响应时间
	ReqBody  any                `json:"reqBody" bson:"reqBody"`   // 请求报文
	RespBody any                `json:"respBody" bson:"respBody"` // 响应报文
}
