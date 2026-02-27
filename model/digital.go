package model

// DigitalInvoiceAuthRequest 数电发票认证请求参数
type DigitalInvoiceAuthRequest struct {
	BaseRequest
	MerchantID         string `json:"merchantId"`         // 银商商户号
	TerminalID         string `json:"terminalId"`         // 银商终端号
	InvoiceNo          string `json:"invoiceNo"`          // 数电发票号码
	InvoiceCode        string `json:"invoiceCode"`        // 数电发票代码
	InvoiceDate        string `json:"invoiceDate"`        // 开票日期（yyyyMMdd）
	AuthType           string `json:"authType"`           // 认证类型（IN-进项/OUT-销项）
}

// DigitalInvoiceAuthResponse 数电发票认证响应参数
type DigitalInvoiceAuthResponse struct {
	BaseResponse
	InvoiceNo          string `json:"invoiceNo"`          // 数电发票号码
	InvoiceCode        string `json:"invoiceCode"`        // 数电发票代码
	AuthStatus         string `json:"authStatus"`         // 认证状态（SUCCESS-成功/FAIL-失败）
	AuthTime           string `json:"authTime"`           // 认证时间
	AuthResult         string `json:"authResult,omitempty"`// 认证结果描述
}

// DigitalInvoiceDownloadRequest 数电发票下载请求参数
type DigitalInvoiceDownloadRequest struct {
	BaseRequest
	MerchantID         string `json:"merchantId"`         // 银商商户号
	TerminalID         string `json:"terminalId"`         // 银商终端号
	InvoiceNo          string `json:"invoiceNo"`          // 数电发票号码
	InvoiceCode        string `json:"invoiceCode"`        // 数电发票代码
	DownloadType       string `json:"downloadType"`       // 下载类型（PDF/XML）
}

// DigitalInvoiceDownloadResponse 数电发票下载响应参数
type DigitalInvoiceDownloadResponse struct {
	BaseResponse
	DownloadURL        string `json:"downloadUrl"`        // 下载地址
	ExpireTime         string `json:"expireTime"`         // 下载地址过期时间
	FileSize           int64  `json:"fileSize"`           // 文件大小（字节）
}