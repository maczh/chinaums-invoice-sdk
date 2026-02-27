package client

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/maczh/chinaums-invoice-sdk/util"
	"github.com/sirupsen/logrus"
)

// Client 银联商务发票服务客户端
type Client struct {
	restyClient *resty.Client
	config      *Config
}

var logger = logrus.New()

var client *Client

// Config 客户端配置
type Config struct {
	BaseURL    string // 接口基础地址（测试/生产）
	MsgSrc     string // 消息来源
	SignKey    string // 签名密钥
	MerchantID string // 银商商户号
	TerminalID string // 银商终端号
	Timeout    time.Duration
	IsProd     bool // 是否生产环境
}

// NewClient 创建客户端
func NewClient(env string, opts ...Option) (*Client, error) {
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

	client := &Client{
		config: &config,
		restyClient: resty.New().
			SetBaseURL(config.BaseURL).
			SetTimeout(config.Timeout).
			SetHeader("Content-Type", "application/json;charset=UTF-8"),
	}

	// 应用选项
	for _, opt := range opts {
		opt(client)
	}

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

	// 生成签名
	reqMap["sign"] = util.Sign(reqMap, client.config.SignKey)

	// 输出请求日志
	logger.Debugf("Request: %s", util.ToJsonString(reqMap))

	// 发送请求
	_, err = client.restyClient.R().
		SetBody(reqMap).
		SetResult(resp).
		Post(client.restyClient.BaseURL)
	if err != nil {
		return err
	}

	// 输出响应日志
	logger.Debugf("Response: %s", util.ToJsonString(resp))

	// 校验响应签名
	respMap, _ := util.StructToMap(resp)
	sign := respMap["sign"].(string)
	delete(respMap, "sign")
	if util.Sign(respMap, client.config.SignKey) != sign {
		return errors.New("response sign verify failed")
	}

	// 检查响应状态
	if respMap["resultCode"].(string) != "SUCCESS" {
		return errors.New(respMap["resultMsg"].(string))
	}

	return nil
}
