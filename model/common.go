package model

// BaseRequest 公共请求参数
type BaseRequest struct {
	MsgID            string `json:"msgId"`		// 消息ID（唯一）
	MsgSrc           string `json:"msgSrc"`
	MsgType          string `json:"msgType"`	//业务类型
	RequestTimestamp any `json:"requestTimestamp"`	// 请求时间戳（yyyy-MM-dd HH:mm:ss）
	SrcReserve       string `json:"srcReserve,omitempty"`
}

func (r *BaseRequest) SetMsgID(id string)            { r.MsgID = id }
func (r *BaseRequest) SetMsgSrc(src string)          { r.MsgSrc = src }
func (r *BaseRequest) SetMsgType(t string)           { r.MsgType = t }
func (r *BaseRequest) SetRequestTimestamp(ts string) { r.RequestTimestamp = ts }
func (r *BaseRequest) SetSrcReserve(reserve string)  { r.SrcReserve = reserve }

// BaseResponse 公共响应参数
type BaseResponse struct {
	MsgID             string `json:"msgId"`
	MsgSrc            string `json:"msgSrc"`
	MsgType           string `json:"msgType"`
	ResultCode        string `json:"resultCode"`		// 结果码（SUCCESS-成功）
	ResultMsg         string `json:"resultMsg,omitempty"`
	ResponseTimestamp string `json:"responseTimestamp"`	// 响应时间戳（yyyy-MM-dd HH:mm:ss）
	SrcReserve        string `json:"srcReserve,omitempty"`
	Sign              string `json:"sign"`
}

func (r *BaseResponse) GetResultCode() string { return r.ResultCode }
func (r *BaseResponse) GetResultMsg() string  { return r.ResultMsg }
