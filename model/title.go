package model

// FuzzyQueryInvoiceTitleRequest 抬头模糊查询请求参数
type FuzzyQueryInvoiceTitleRequest struct {
	BaseRequest
	Name    string `json:"name"`    // 抬头名称
	TaxCode string `json:"taxCode"` // 税号
}

// FuzzyQueryInvoiceTitleResponse 抬头模糊查询响应参数
type FuzzyQueryInvoiceTitleResponse struct {
	BaseResponse
	TitleList []HeadListItem `json:"titleList"` // 抬头列表
}

// HeadListItem 抬头列表项
type HeadListItem struct {
	Name      string `json:"name"`                // 抬头名称
	TaxCode   string `json:"taxCode"`             // 税号
	Address   string `json:"address,omitempty"`   // 地址
	Telephone string `json:"telephone,omitempty"` // 电话
	Bank      string `json:"bank,omitempty"`      // 开户行
	Account   string `json:"account,omitempty"`   // 银行账号
}

// WechatRecognizeHeadRequest 微信抬头识别请求参数
type WechatRecognizeHeadRequest struct {
	BaseRequest
	MerchantID string `json:"merchantId"` // 银商商户号
	TerminalID string `json:"terminalId"` // 银商终端号
	AuthCode   string `json:"authCode"`   // 微信授权码
}

// WechatRecognizeHeadResponse 微信抬头识别响应参数
type WechatRecognizeHeadResponse struct {
	BaseResponse
	HeadName    string `json:"headName"`              // 抬头名称
	TaxCode     string `json:"taxCode"`               // 税号
	Address     string `json:"address,omitempty"`     // 地址
	Phone       string `json:"phone,omitempty"`       // 电话
	BankName    string `json:"bankName,omitempty"`    // 开户行
	BankAccount string `json:"bankAccount,omitempty"` // 银行账号
}
