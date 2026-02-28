package model

// DirectIssueRequest 直接开票请求参数
type DirectIssueRequest struct {
	BaseRequest
	InvoiceMaterial       string `json:"invoiceMaterial"`                 // 发票材质: 纸质发票：PAPER 电子发票：ELECTRONIC
	InvoiceType           string `json:"invoiceType"`                     // 发票类型: 普通发票：PLAIN 增值税专用发票：VAT
	MerchantId            string `json:"merchantId"`                      // 银商商户号(15位)
	TerminalId            string `json:"terminalId"`                      // 银商终端号(8位)
	MerOrderDate          string `json:"merOrderDate"`                    // 开票订单日期(格式yyyyMMdd)
	MerOrderId            string `json:"merOrderId"`                      // 开票订单号(1-32位)
	BuyerName             string `json:"buyerName,omitempty"`             // 买方名称(1-128字符)
	BuyerTaxCode          string `json:"buyerTaxCode,omitempty"`          // 买方纳税人识别号(1-32字符)
	BuyerAddress          string `json:"buyerAddress,omitempty"`          // 买方地址(1-128字符)
	BuyerTelephone        string `json:"buyerTelephone,omitempty"`        // 买方电话(1-16字符)
	BuyerBank             string `json:"buyerBank,omitempty"`             // 买方开户行(1-64字符)
	BuyerAccount          string `json:"buyerAccount,omitempty"`          // 买方银行账号(1-32字符)
	BuyerBrevityCode      string `json:"buyerBrevityCode,omitempty"`      // 买方简码(1-32字符)
	BuyerGroupBrevityCode string `json:"buyerGroupBrevityCode,omitempty"` // 买方集团简码(1-32字符)
	Amount                int64  `json:"amount"`                          // 含税总金额(单位为分)
	DeductionAmount       int64  `json:"deductionAmount"`                 // 扣除额(单位为分)
	GoodsDetail           string `json:"goodsDetail,omitempty"`           // 商品明细(JSON字符串，最大15000字符)
	Remark                string `json:"remark,omitempty"`                // 备注(1-128字符)
	NotifyMobileNo        string `json:"notifyMobileNo,omitempty"`        // 消费者手机号，用于短信通知电票(1-16字符)
	NotifyEmail           string `json:"notifyEmail,omitempty"`           // 消费者邮箱，用于邮件通知电票(1-64字符)
	NotifyUrl             string `json:"notifyUrl,omitempty"`             // 商户通知地址，用于开票结果通知(1-256字符)
	MerWxAppId            string `json:"merWxAppId,omitempty"`            // 商户微信AppId，用于微信插卡(1-32字符)
	MerWxOrderId          string `json:"merWxOrderId,omitempty"`          // 商户微信OrderId，用于微信插卡(1-64字符)
	AlipayUserId          string `json:"alipayUserId,omitempty"`          // 用户支付宝userId，用于支付宝插卡(1-16字符)
	StoreId               string `json:"storeId,omitempty"`               // 门店号(1-32字符)
	StoreName             string `json:"storeName,omitempty"`             // 门店名称(1-128字符)
	Payee                 string `json:"payee,omitempty"`                 // 收款人(1-8字符)
	Checker               string `json:"checker,omitempty"`               // 复核人(1-8字符)
	Drawer                string `json:"drawer,omitempty"`                // 开票人(1-8字符)
	OperateId             string `json:"operateId,omitempty"`             // 营运ID(1-18字符)
	VehicleNo             string `json:"vehicleNo,omitempty"`             // 车牌号(1-10字符)
	OutTradeNo            string `json:"outTradeNo,omitempty"`            // 三方自定义的订单号(1-64字符)
	NotFaceRemark         string `json:"notFaceRemark,omitempty"`         // 非票面备注(1-256字符)
	SpecialInvoiceFlag    string `json:"specialInvoiceFlag,omitempty"`    // 特殊票种标识(2字符)
	InvoiceSpecialInfo    string `json:"invoiceSpecialInfo,omitempty"`    // 特殊票种信息(JSON字符串，最大15000字符)
	ConfirmIssue          string `json:"confirmIssue,omitempty"`          // 全电成品油单价过低是否确认开具标识(2字符)
	InvoiceBalanceInfo    string `json:"invoiceBalanceInfo,omitempty"`    // 数电发票差额扣除额凭证明细(JSON数组)
}

// DirectIssueResponse 直接开票响应参数
type DirectIssueResponse struct {
	BaseResponse
	Status                 string        `json:"status"`                          // 发票状态: ISSUING/ISSUED/CLOSED/SPLITED
	InvoiceMaterial        string        `json:"invoiceMaterial"`                 // 发票材质: PAPER/ELECTRONIC
	InvoiceType            string        `json:"invoiceType"`                     // 发票类型: PLAIN/VAT
	InvoiceNo              string        `json:"invoiceNo"`                       // 发票号码(1-20字符)
	InvoiceCode            string        `json:"invoiceCode"`                     // 发票代码(1-20字符)
	CheckCode              string        `json:"checkCode"`                       // 校验码(1-32字符)
	CipherCode             string        `json:"cipherCode"`                      // 密码区(1-256字符)
	IssueDate              string        `json:"issueDate"`                       // 开票日期(yyyy-MM-dd HH:mm:ss格式)
	DeviceNo               string        `json:"deviceNo,omitempty"`              // 机器编号(1-32字符)
	QRCodeID               string        `json:"qrCodeId,omitempty"`              // 二维码唯一id(40字符)
	StoreID                string        `json:"storeId,omitempty"`               // 门店编号(1-32字符)
	StoreName              string        `json:"storeName,omitempty"`             // 门店名称(1-128字符)
	MerchantName           string        `json:"merchantName,omitempty"`          // 银商商户名称(1-128字符)
	MerchantID             string        `json:"merchantId,omitempty"`            // 银商商户号(15字符)
	TerminalID             string        `json:"terminalId,omitempty"`            // 银商终端号(8字符)
	MerOrderID             string        `json:"merOrderId,omitempty"`            // 商户订单号(1-32字符)
	MerOrderDate           string        `json:"merOrderDate,omitempty"`          // 商户订单日期(yyyyMMdd格式)
	BuyerName              string        `json:"buyerName,omitempty"`             // 买方名称(1-128字符)
	BuyerTaxCode           string        `json:"buyerTaxCode,omitempty"`          // 买方纳税人识别号(1-32字符)
	BuyerAddress           string        `json:"buyerAddress,omitempty"`          // 买方地址(1-128字符)
	BuyerTelephone         string        `json:"buyerTelephone,omitempty"`        // 买方电话(1-16字符)
	BuyerBank              string        `json:"buyerBank,omitempty"`             // 买方开户行(1-64字符)
	BuyerAccount           string        `json:"buyerAccount,omitempty"`          // 买方银行账号(1-32字符)
	BuyerBrevityCode       string        `json:"buyerBrevityCode,omitempty"`      // 买方简码(1-32字符)
	BuyerGroupBrevityCode  string        `json:"buyerGroupBrevityCode,omitempty"` // 买方集团简码(1-32字符)
	SellerTaxCode          string        `json:"sellerTaxCode,omitempty"`         // 卖方纳税人识别号(1-32字符)
	SellerAddress          string        `json:"sellerAddress,omitempty"`         // 卖方地址(1-128字符)
	SellerTelephone        string        `json:"sellerTelephone,omitempty"`       // 卖方电话(1-16字符)
	SellerBank             string        `json:"sellerBank,omitempty"`            // 卖方开户行(1-64字符)
	SellerAccount          string        `json:"sellerAccount,omitempty"`         // 卖方账号(1-32字符)
	Payee                  string        `json:"payee,omitempty"`                 // 收款人(1-8字符)
	Checker                string        `json:"checker,omitempty"`               // 复核人(1-8字符)
	Drawer                 string        `json:"drawer,omitempty"`                // 开票人(1-8字符)
	Remark                 string        `json:"remark,omitempty"`                // 备注(1-128字符)
	OutTradeNo             string        `json:"outTradeNo,omitempty"`            // 三方自定义的订单号(1-64字符)
	NotFaceRemark          string        `json:"notFaceRemark,omitempty"`         // 非票面备注(1-256字符)
	DeductionAmount        float64       `json:"deductionAmount"`                 // 扣除额(单位为元)
	TotalPriceIncludingTax float64       `json:"totalPriceIncludingTax"`          // 含税总金额(单位为元)
	TotalTax               float64       `json:"totalTax"`                        // 税额(单位为元)
	TotalPrice             float64       `json:"totalPrice"`                      // 不含税总金额(单位为元)
	GoodsDetail            string        `json:"goodsDetail,omitempty"`           // 商品明细(1-1024字符)
	NotifyMobileNo         string        `json:"notifyMobileNo,omitempty"`        // 消费者手机号(1-16字符)
	NotifyEmail            string        `json:"notifyEmail,omitempty"`           // 消费者邮箱(1-64字符)
	NotifyUrl              string        `json:"notifyUrl,omitempty"`             // 商户通知地址(1-128字符)
	NotifyMerEmail         string        `json:"notifyMerEmail,omitempty"`        // 商户邮箱(1-64字符)
	PdfUrl                 string        `json:"pdfUrl,omitempty"`                // PDF下载链接(1-128字符)
	PdfPreviewUrl          string        `json:"pdfPreviewUrl,omitempty"`         // PDF预览链接(1-128字符)
	OfdUrl                 string        `json:"ofdUrl,omitempty"`                // OFD下载链接(1-128字符)
	OfdPreviewUrl          string        `json:"ofdPreviewUrl,omitempty"`         // OFD预览链接(1-128字符)
	XmlUrl                 string        `json:"xmlUrl,omitempty"`                // XML文件地址(1-128字符)
	XmlPreviewUrl          string        `json:"xmlPreviewUrl,omitempty"`         // XML预览链接(1-128字符)
	ErrMsg                 string        `json:"errMsg,omitempty"`                // 开票结果信息(1-256字符)
	SubList                []interface{} `json:"subList"`                         // 子订单信息(JsonArray格式的字符串)
}

// AsyncIssueRequest 异步开票请求参数（同直接开票，仅msgType不同）
type AsyncIssueRequest DirectIssueRequest

// AsyncIssueResponse 异步开票响应参数
type AsyncIssueResponse DirectIssueResponse

// GoodsDetail 商品明细结构体（序列化后传入goodsDetail字段）
type GoodsDetail struct {
	Index                 int     `json:"index"`                      // 行号（从1开始）
	Attribute             string  `json:"attribute"`                  // 商品属性（0-正常/1-折扣/2-不征税）
	Name                  string  `json:"name"`                       // 商品名称
	SN                    string  `json:"sn"`                         // 商品税目编码
	Model                 string  `json:"model,omitempty"`            // 规格型号
	Unit                  string  `json:"unit,omitempty"`             // 单位
	Quantity              float64 `json:"quantity"`                   // 数量
	PriceIncludingTax     float64 `json:"priceIncludingTax"`          // 含税单价（元）
	TaxRate               float64 `json:"taxRate"`                    // 税率（如6/9/13,百分比）
	FreeTaxType           string  `json:"freeTaxType,omitempty"`      // 免税类型（0-普通免税/1-特殊免税）
	UnitPriceIncludingTax float64 `json:"unitPriceIncludingTax"`      // 含税单价（元）
	PreferPolicyFlag      string  `json:"preferPolicyFlag,omitempty"` // 优先政策标志（0-普通发票/1-专票）
	VatSpecial            string  `json:"vatSpecial,omitempty"`       // 增值税特殊处理标志（0-普通发票/1-专票）
}

// IssueCallbackRequest 开票结果通知请求参数
type IssueCallbackRequest struct {
	Status                 string  `json:"status"`                 // 发票状态: ISSUED/REVERSED/INVALIDED/CLOSED/CANCELED/SPLITED
	ErrMsg                 string  `json:"errMsg"`                 // 开票结果信息(1-256字符)
	MerchantName           string  `json:"merchantName"`           // 银商商户名称(1-128字符)
	MerchantId             string  `json:"merchantId"`             // 银商商户号(15字符)
	TerminalId             string  `json:"terminalId"`             // 银商终端号(8字符)
	MerOrderId             string  `json:"merOrderId"`             // 商户订单号(1-32字符)
	MerOrderDate           string  `json:"merOrderDate"`           // 商户订单日期(yyyyMMdd格式)
	QRCodeID               string  `json:"qrCodeId"`               // 二维码唯一id(40字符)
	InvoiceMaterial        string  `json:"invoiceMaterial"`        // 发票材质: PAPER/ELECTRONIC
	InvoiceType            string  `json:"invoiceType"`            // 发票类型: PLAIN/VAT
	InvoiceNo              string  `json:"invoiceNo"`              // 发票号码(1-20字符)
	InvoiceCode            string  `json:"invoiceCode"`            // 发票代码(1-20字符)
	CheckCode              string  `json:"checkCode"`              // 校验码(1-32字符)
	CipherCode             string  `json:"cipherCode"`             // 密码区(1-256字符)
	IssueDate              string  `json:"issueDate"`              // 开票日期(yyyy-MM-dd HH:mm:ss格式)
	ReverseInvoiceNo       string  `json:"reverseInvoiceNo"`       // 红票的发票号码(1-20字符)
	ReverseInvoiceCode     string  `json:"reverseInvoiceCode"`     // 红票的发票代码(1-20字符)
	ReverseCheckCode       string  `json:"reverseCheckCode"`       // 红票的校验码(1-32字符)
	ReverseCipherCode      string  `json:"reverseCipherCode"`      // 红票的密码区(1-256字符)
	ReverseDate            string  `json:"reverseDate"`            // 红冲日期(yyyy-MM-dd HH:mm:ss格式)
	DeviceNo               string  `json:"deviceNo"`               // 机器编号(1-32字符)
	StoreId                string  `json:"storeId"`                // 门店编号(1-32字符)
	StoreName              string  `json:"storeName"`              // 门店名称(1-128字符)
	BuyerName              string  `json:"buyerName"`              // 买方名称(1-128字符)
	BuyerTaxCode           string  `json:"buyerTaxCode"`           // 买方纳税人识别号(1-32字符)
	BuyerAddress           string  `json:"buyerAddress"`           // 买方地址(1-128字符)
	BuyerTelephone         string  `json:"buyerTelephone"`         // 买方电话(1-16字符)
	BuyerBank              string  `json:"buyerBank"`              // 买方开户行(1-64字符)
	BuyerAccount           string  `json:"buyerAccount"`           // 买方银行账号(1-32字符)
	SellerName             string  `json:"sellerName"`             // 卖方名称(1-128字符)
	SellerTaxCode          string  `json:"sellerTaxCode"`          // 卖方纳税人识别号(1-32字符)
	SellerAddress          string  `json:"sellerAddress"`          // 卖方地址(1-128字符)
	SellerTelephone        string  `json:"sellerTelephone"`        // 卖方电话(1-16字符)
	SellerBank             string  `json:"sellerBank"`             // 卖方开户行(1-64字符)
	SellerAccount          string  `json:"sellerAccount"`          // 卖方账号(1-32字符)
	Payee                  string  `json:"payee"`                  // 收款人(1-8字符)
	Checker                string  `json:"checker"`                // 复核人(1-8字符)
	Drawer                 string  `json:"drawer"`                 // 开票人(1-8字符)
	Remark                 string  `json:"remark"`                 // 备注(1-128字符)
	TaxMethod              string  `json:"taxMethod"`              // 征税方式: NORMAL/DEDUCTION
	DeductionAmount        float64 `json:"deductionAmount"`        // 扣除额(单位为元)
	TotalPriceIncludingTax float64 `json:"totalPriceIncludingTax"` // 含税总金额(单位为元)
	TotalTax               float64 `json:"totalTax"`               // 税额(单位为元)
	TotalPrice             float64 `json:"totalPrice"`             // 不含税总金额(单位为元)
	GoodsDetail            string  `json:"goodsDetail"`            // 商品明细(1-1024字符)
	NotifyMobileNo         string  `json:"notifyMobileNo"`         // 消费者手机号(1-16字符)
	NotifyEmail            string  `json:"notifyEmail"`            // 消费者邮箱(1-64字符)
	NotifyUrl              string  `json:"notifyUrl"`              // 商户通知地址(1-128字符)
	NotifyMerEmail         string  `json:"notifyMerEmail"`         // 商户邮箱(1-64字符)
	MerWxAppId             string  `json:"merWxAppId"`             // 商户微信AppId(1-32字符)
	MerWxOrderId           string  `json:"merWxOrderId"`           // 商户微信OrderId(1-64字符)
	PdfUrl                 string  `json:"pdfUrl"`                 // PDF下载链接(1-128字符)
	PdfPreviewUrl          string  `json:"pdfPreviewUrl"`          // PDF预览链接(1-128字符)
	Source                 string  `json:"source"`                 // 来源系统(同msgSrc)
	Sign                   string  `json:"sign"`                   // 签名值
	SubList                []any   `json:"subList"`                // 子订单信息(JsonArray格式的字符串)
}

// IssueCallbackResponse 发票开具回调响应参数
type IssueCallbackResponse struct {
	ResultCode string `json:"resultCode"` // 结果码: SUCCESS/FAIL
	ResultMsg  string `json:"resultMsg"`  // 结果信息
}
