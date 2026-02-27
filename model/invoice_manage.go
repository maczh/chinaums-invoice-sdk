package model

// QueryInvoiceStatusRequest 发票状态查询请求参数
type QueryInvoiceStatusRequest struct {
	BaseRequest
	MerchantID         string `json:"merchantId"`         // 银商商户号
	TerminalID         string `json:"terminalId"`         // 银商终端号
	MerOrderID         string `json:"merOrderId,omitempty"`// 商户订单号（二选一）
	MerOrderDate       string `json:"merOrderDate"`       // 订单日期（yyyyMMdd）
}

// QueryInvoiceStatusResponse 发票状态查询响应参数
type QueryInvoiceStatusResponse struct {
	BaseResponse
    Status               string `json:"status"`                 // 发票状态: PENDING/CLOSED/CANCELED/ISSUING/ISSUED/REVERSING/REVERSED/INVALIDING/INVALIDED/SPLITED
    InvoiceMaterial      string `json:"invoiceMaterial"`        // 发票材质: PAPER/ELECTRONIC
    InvoiceType          string `json:"invoiceType"`            // 发票类型: PLAIN/VAT
    InvoiceNo            string `json:"invoiceNo"`              // 发票号码(1-20字符)
    InvoiceCode          string `json:"invoiceCode"`            // 发票代码(1-20字符)
    CheckCode            string `json:"checkCode"`              // 校验码(1-32字符)
    CipherCode           string `json:"cipherCode"`             // 密码区(1-256字符)
    IssueDate            string `json:"issueDate"`              // 开票日期(yyyy-MM-dd HH:mm:ss格式)
    ReverseInvoiceNo     string `json:"reverseInvoiceNo"`       // 红票的发票号码(1-20字符)
    ReverseInvoiceCode   string `json:"reverseInvoiceCode"`     // 红票的发票代码(1-20字符)
    ReverseCheckCode     string `json:"reverseCheckCode"`       // 红票的校验码(1-32字符)
    ReverseCipherCode    string `json:"reverseCipherCode"`      // 红票的密码区(1-256字符)
    ReverseDate          string `json:"reverseDate"`            // 红冲日期(yyyy-MM-dd HH:mm:ss格式)
    DeviceNo             string `json:"deviceNo"`               // 机器编号(1-32字符)
    QRCodeID             string `json:"qrCodeId"`               // 二维码唯一id(40字符)
    StoreID              string `json:"storeId"`                // 门店编号(1-32字符)
    StoreName            string `json:"storeName"`              // 门店名称(1-128字符)
    MerchantName         string `json:"merchantName"`           // 银商商户名称(1-128字符)
    MerchantID           string `json:"merchantId"`             // 银商商户号(15字符)
    TerminalID           string `json:"terminalId"`             // 银商终端号(8字符)
    MerOrderID           string `json:"merOrderId"`             // 商户订单号(1-32字符)
    MerOrderDate         string `json:"merOrderDate"`           // 商户订单日期(yyyyMMdd格式)
    BuyerName            string `json:"buyerName"`              // 买方名称(1-128字符)
    BuyerTaxCode         string `json:"buyerTaxCode"`           // 买方纳税人识别号(1-32字符)
    BuyerAddress         string `json:"buyerAddress"`           // 买方地址(1-128字符)
    BuyerTelephone       string `json:"buyerTelephone"`         // 买方电话(1-16字符)
    BuyerBank            string `json:"buyerBank"`              // 买方开户行(1-64字符)
    BuyerAccount         string `json:"buyerAccount"`           // 买方银行账号(1-32字符)
    SellerName           string `json:"sellerName"`             // 卖方名称(1-128字符)
    SellerTaxCode        string `json:"sellerTaxCode"`          // 卖方纳税人识别号(1-32字符)
    SellerAddress        string `json:"sellerAddress"`          // 卖方地址(1-128字符)
    SellerTelephone      string `json:"sellerTelephone"`        // 卖方电话(1-16字符)
    SellerBank           string `json:"sellerBank"`             // 卖方开户行(1-64字符)
    SellerAccount        string `json:"sellerAccount"`          // 卖方账号(1-32字符)
    Payee                string `json:"payee"`                  // 收款人(1-8字符)
    Checker              string `json:"checker"`                // 复核人(1-8字符)
    Drawer               string `json:"drawer"`                 // 开票人(1-8字符)
    Remark               string `json:"remark"`                 // 备注(1-128字符)
    TaxMethod            string `json:"taxMethod"`              // 征税方式: NORMAL/DEDUCTION
    DeductionAmount      string `json:"deductionAmount"`        // 扣除额(单位为元)
    TotalPriceIncludingTax string `json:"totalPriceIncludingTax"` // 含税总金额(单位为元)
    TotalTax             string `json:"totalTax"`               // 税额(单位为元)
    TotalPrice           string `json:"totalPrice"`             // 不含税总金额(单位为元)
    GoodsDetail          string `json:"goodsDetail"`            // 商品明细(1-1024字符)
    NotifyMobileNo       string `json:"notifyMobileNo"`         // 消费者手机号(1-16字符)
    NotifyEmail          string `json:"notifyEmail"`            // 消费者邮箱(1-64字符)
    NotifyUrl            string `json:"notifyUrl"`              // 商户通知地址(1-128字符)
    NotifyMerEmail       string `json:"notifyMerEmail"`         // 商户邮箱(1-64字符)
    MerWxAppID           string `json:"merWxAppId"`             // 商户微信AppId(1-32字符)
    MerWxOrderID         string `json:"merWxOrderId"`           // 商户微信OrderId(1-64字符)
    IssueQRCode          string `json:"issueQrCode"`            // 开票二维码(1-128字符)
    PdfUrl               string `json:"pdfUrl"`                 // PDF下载链接(1-128字符)
    PdfPreviewUrl        string `json:"pdfPreviewUrl"`          // PDF预览链接(1-128字符)
    ErrMsg               string `json:"errMsg"`                 // 开票结果信息(1-256字符)
    SubList              string `json:"subList"`                // 子订单信息(JsonArray格式的字符串)
}


// RedReverseInvoiceRequest 红冲发票请求参数
type RedReverseInvoiceRequest struct {
	BaseRequest
	MerchantID         string `json:"merchantId"`         // 银商商户号
	TerminalID         string `json:"terminalId"`         // 银商终端号
	MerOrderID         string `json:"merOrderId"`         // 红冲订单号（唯一）
	MerOrderDate       string `json:"merOrderDate"`       // 订单日期（yyyyMMdd）
	RedNotificationNo  string `json:"redNotificationNo,omitempty"`// 红字信息表编号
}

// RedReverseInvoiceResponse 红冲发票响应参数
type RedReverseInvoiceResponse struct {
    Status             string `json:"status"`              // 发票状态: REVERSING/REVERSED/CLOSED
    InvoiceMaterial    string `json:"invoiceMaterial"`     // 发票材质: PAPER/ELECTRONIC
    InvoiceType        string `json:"invoiceType"`         // 发票类型: PLAIN/VAT
    InvoiceNo          string `json:"invoiceNo"`           // 发票号码(1-20字符)
    InvoiceCode        string `json:"invoiceCode"`         // 发票代码(1-20字符)
    CheckCode          string `json:"checkCode"`           // 校验码(1-32字符)
    CipherCode         string `json:"cipherCode"`          // 密码区(1-256字符)
    IssueDate          string `json:"issueDate"`           // 开票日期(yyyy-MM-dd HH:mm:ss格式)
    ReverseInvoiceNo   string `json:"reverseInvoiceNo"`    // 红票的发票号码(1-20字符)
    ReverseInvoiceCode string `json:"reverseInvoiceCode"`  // 红票的发票代码(1-20字符)
    ReverseCheckCode   string `json:"reverseCheckCode"`    // 红票的校验码(1-32字符)
    ReverseCipherCode  string `json:"reverseCipherCode"`   // 红票的密码区(1-256字符)
    ReverseDate        string `json:"reverseDate"`         // 红冲日期(yyyy-MM-dd HH:mm:ss格式)
    DeviceNo           string `json:"deviceNo"`            // 机器编号(1-32字符)
    QRCodeID           string `json:"qrCodeId"`            // 二维码唯一id(40字符)
    StoreID            string `json:"storeId"`             // 门店编号(1-32字符)
    StoreName          string `json:"storeName"`           // 门店名称(1-128字符)
    MerchantName       string `json:"merchantName"`        // 银商商户名称(1-128字符)
    MerchantID         string `json:"merchantId"`          // 银商商户号(15字符)
    TerminalID         string `json:"terminalId"`          // 银商终端号(8字符)
    MerOrderID         string `json:"merOrderId"`          // 商户订单号(1-32字符)
    MerOrderDate       string `json:"merOrderDate"`        // 商户订单日期(yyyyMMdd格式)
    BuyerName          string `json:"buyerName"`           // 买方名称(1-128字符)
    BuyerTaxCode       string `json:"buyerTaxCode"`        // 买方纳税人识别号(1-32字符)
    BuyerAddress       string `json:"buyerAddress"`        // 买方地址(1-128字符)
    BuyerTelephone     string `json:"buyerTelephone"`      // 买方电话(1-16字符)
    BuyerBank          string `json:"buyerBank"`           // 买方开户行(1-64字符)
    BuyerAccount       string `json:"buyerAccount"`        // 买方银行账号(1-32字符)
    SellerName         string `json:"sellerName"`          // 卖方名称(1-128字符)
    SellerTaxCode      string `json:"sellerTaxCode"`       // 卖方纳税人识别号(1-32字符)
    SellerAddress      string `json:"sellerAddress"`       // 卖方地址(1-128字符)
    SellerTelephone    string `json:"sellerTelephone"`     // 卖方电话(1-16字符)
    SellerBank         string `json:"sellerBank"`          // 卖方开户行(1-64字符)
    SellerAccount      string `json:"sellerAccount"`       // 卖方账号(1-32字符)
    Payee              string `json:"payee"`               // 收款人(1-8字符)
    Checker            string `json:"checker"`             // 复核人(1-8字符)
    Drawer             string `json:"drawer"`              // 开票人(1-8字符)
    Remark             string `json:"remark"`              // 备注(1-128字符)
    TaxMethod          string `json:"taxMethod"`           // 征税方式: NORMAL/DEDUCTION
    DeductionAmount    string `json:"deductionAmount"`     // 扣除额(单位为元)
    TotalPriceIncludingTax string `json:"totalPriceIncludingTax"` // 含税总金额(单位为元)
    TotalTax           string `json:"totalTax"`            // 税额(单位为元)
    TotalPrice         string `json:"totalPrice"`          // 不含税总金额(单位为元)
    GoodsDetail        string `json:"goodsDetail"`         // 商品明细(1-1024字符)
    NotifyMobileNo     string `json:"notifyMobileNo"`      // 消费者手机号(1-16字符)
    NotifyEmail        string `json:"notifyEmail"`         // 消费者邮箱(1-64字符)
    NotifyUrl          string `json:"notifyUrl"`           // 商户通知地址(1-128字符)
    NotifyMerEmail     string `json:"notifyMerEmail"`      // 商户邮箱(1-64字符)
    PdfUrl             string `json:"pdfUrl"`              // PDF下载链接(1-128字符)
    PdfPreviewUrl      string `json:"pdfPreviewUrl"`       // PDF预览链接(1-128字符)
    ErrMsg             string `json:"errMsg"`              // 开票结果信息(1-256字符)
    SubList            string `json:"subList"`             // 子订单信息(JsonArray格式的字符串)
}

// InvalidInvoiceRequest 作废发票请求参数
type InvalidInvoiceRequest struct {
	BaseRequest
	MerchantID         string `json:"merchantId"`         // 银商商户号
	TerminalID         string `json:"terminalId"`         // 银商终端号
	MerOrderID         string `json:"merOrderId"`         // 红冲订单号（唯一）
	MerOrderDate       string `json:"merOrderDate"`       // 订单日期（yyyyMMdd）
	InvalidReason      string `json:"invalidReason"`      // 作废原因
}

// InvalidInvoiceResponse 作废发票响应参数
type InvalidInvoiceResponse struct {
	BaseResponse
    Status             string `json:"status"`              // 发票状态: INVALIDING/INVALIDED/CLOSED
    InvoiceMaterial    string `json:"invoiceMaterial"`     // 发票材质: PAPER/ELECTRONIC
    InvoiceType        string `json:"invoiceType"`         // 发票类型: PLAIN/VAT
    InvoiceNo          string `json:"invoiceNo"`           // 发票号码(1-20字符)
    InvoiceCode        string `json:"invoiceCode"`         // 发票代码(1-20字符)
    CheckCode          string `json:"checkCode"`           // 校验码(1-32字符)
    CipherCode         string `json:"cipherCode"`          // 密码区(1-256字符)
    IssueDate          string `json:"issueDate"`           // 开票日期(yyyy-MM-dd HH:mm:ss格式)
    InvalidDate        string `json:"invalidDate"`         // 作废日期(yyyy-MM-dd HH:mm:ss格式)
    DeviceNo           string `json:"deviceNo"`            // 机器编号(1-32字符)
    QRCodeID           string `json:"qrCodeId"`            // 二维码唯一id(40字符)
    StoreID            string `json:"storeId"`             // 门店编号(1-32字符)
    StoreName          string `json:"storeName"`           // 门店名称(1-128字符)
    MerchantName       string `json:"merchantName"`        // 银商商户名称(1-128字符)
    MerchantID         string `json:"merchantId"`          // 银商商户号(15字符)
    TerminalID         string `json:"terminalId"`          // 银商终端号(8字符)
    MerOrderID         string `json:"merOrderId"`          // 商户订单号(1-32字符)
    MerOrderDate       string `json:"merOrderDate"`        // 商户订单日期(yyyyMMdd格式)
    BuyerName          string `json:"buyerName"`           // 买方名称(1-128字符)
    BuyerTaxCode       string `json:"buyerTaxCode"`        // 买方纳税人识别号(1-32字符)
    BuyerAddress       string `json:"buyerAddress"`        // 买方地址(1-128字符)
    BuyerTelephone     string `json:"buyerTelephone"`      // 买方电话(1-16字符)
    BuyerBank          string `json:"buyerBank"`           // 买方开户行(1-64字符)
    BuyerAccount       string `json:"buyerAccount"`        // 买方银行账号(1-32字符)
    SellerName         string `json:"sellerName"`          // 卖方名称(1-128字符)
    SellerTaxCode      string `json:"sellerTaxCode"`       // 卖方纳税人识别号(1-32字符)
    SellerAddress      string `json:"sellerAddress"`       // 卖方地址(1-128字符)
    SellerTelephone    string `json:"sellerTelephone"`     // 卖方电话(1-16字符)
    SellerBank         string `json:"sellerBank"`          // 卖方开户行(1-64字符)
    SellerAccount      string `json:"sellerAccount"`       // 卖方账号(1-32字符)
    Payee              string `json:"payee"`               // 收款人(1-8字符)
    Checker            string `json:"checker"`             // 复核人(1-8字符)
    Drawer             string `json:"drawer"`              // 开票人(1-8字符)
    Remark             string `json:"remark"`              // 备注(1-128字符)
    TaxMethod          string `json:"taxMethod"`           // 征税方式: NORMAL/DEDUCTION
    DeductionAmount    string `json:"deductionAmount"`     // 扣除额(单位为元)
    TotalPriceIncludingTax string `json:"totalPriceIncludingTax"` // 含税总金额(单位为元)
    TotalTax           string `json:"totalTax"`            // 税额(单位为元)
    TotalPrice         string `json:"totalPrice"`          // 不含税总金额(单位为元)
    GoodsDetail        string `json:"goodsDetail"`         // 商品明细(1-1024字符)
    NotifyMobileNo     string `json:"notifyMobileNo"`      // 消费者手机号(1-16字符)
    NotifyEmail        string `json:"notifyEmail"`         // 消费者邮箱(1-64字符)
    NotifyUrl          string `json:"notifyUrl"`           // 商户通知地址(1-128字符)
    NotifyMerEmail     string `json:"notifyMerEmail"`      // 商户邮箱(1-64字符)
    PdfUrl             string `json:"pdfUrl"`              // PDF下载链接(1-128字符)
    PdfPreviewUrl      string `json:"pdfPreviewUrl"`       // PDF预览链接(1-128字符)
    ErrMsg             string `json:"errMsg"`              // 作废结果信息(1-256字符)
    SubList            string `json:"subList"`             // 子订单信息(JsonArray格式的字符串)
}

// CancelInvoiceRequest 撤消发票请求参数
type CancelInvoiceRequest struct {
	BaseRequest
	MerchantID         string `json:"merchantId"`         // 银商商户号
	TerminalID         string `json:"terminalId"`         // 银商终端号
	MerOrderID         string `json:"merOrderId"`         // 红冲订单号（唯一）
	MerOrderDate       string `json:"merOrderDate"`       // 订单日期（yyyyMMdd）
}

// CancelInvoiceResponse 撤消发票响应参数
type CancelInvoiceResponse struct {
    Status             string `json:"status"`              // 发票状态: CANCELED/CLOSED
    InvoiceMaterial    string `json:"invoiceMaterial"`     // 发票材质: PAPER/ELECTRONIC
    InvoiceType        string `json:"invoiceType"`         // 发票类型: PLAIN/VAT
    InvoiceNo          string `json:"invoiceNo"`           // 发票号码(1-20字符)
    InvoiceCode        string `json:"invoiceCode"`         // 发票代码(1-20字符)
    CheckCode          string `json:"checkCode"`           // 校验码(1-32字符)
    CipherCode         string `json:"cipherCode"`          // 密码区(1-256字符)
    IssueDate          string `json:"issueDate"`           // 开票日期(yyyy-MM-dd HH:mm:ss格式)
    CancelDate         string `json:"cancelDate"`          // 撤销日期(yyyy-MM-dd HH:mm:ss格式)
    DeviceNo           string `json:"deviceNo"`            // 机器编号(1-32字符)
    QRCodeID           string `json:"qrCodeId"`            // 二维码唯一id(40字符)
    StoreID            string `json:"storeId"`             // 门店编号(1-32字符)
    StoreName          string `json:"storeName"`           // 门店名称(1-128字符)
    MerchantName       string `json:"merchantName"`        // 银商商户名称(1-128字符)
    MerchantID         string `json:"merchantId"`          // 银商商户号(15字符)
    TerminalID         string `json:"terminalId"`          // 银商终端号(8字符)
    MerOrderID         string `json:"merOrderId"`          // 商户订单号(1-32字符)
    MerOrderDate       string `json:"merOrderDate"`        // 商户订单日期(yyyyMMdd格式)
    BuyerName          string `json:"buyerName"`           // 买方名称(1-128字符)
    BuyerTaxCode       string `json:"buyerTaxCode"`        // 买方纳税人识别号(1-32字符)
    BuyerAddress       string `json:"buyerAddress"`        // 买方地址(1-128字符)
    BuyerTelephone     string `json:"buyerTelephone"`      // 买方电话(1-16字符)
    BuyerBank          string `json:"buyerBank"`           // 买方开户行(1-64字符)
    BuyerAccount       string `json:"buyerAccount"`        // 买方银行账号(1-32字符)
    SellerName         string `json:"sellerName"`          // 卖方名称(1-128字符)
    SellerTaxCode      string `json:"sellerTaxCode"`       // 卖方纳税人识别号(1-32字符)
    SellerAddress      string `json:"sellerAddress"`       // 卖方地址(1-128字符)
    SellerTelephone    string `json:"sellerTelephone"`     // 卖方电话(1-16字符)
    SellerBank         string `json:"sellerBank"`          // 卖方开户行(1-64字符)
    SellerAccount      string `json:"sellerAccount"`       // 卖方账号(1-32字符)
    Payee              string `json:"payee"`               // 收款人(1-8字符)
    Checker            string `json:"checker"`             // 复核人(1-8字符)
    Drawer             string `json:"drawer"`              // 开票人(1-8字符)
    Remark             string `json:"remark"`              // 备注(1-128字符)
    TaxMethod          string `json:"taxMethod"`           // 征税方式: NORMAL/DEDUCTION
    DeductionAmount    string `json:"deductionAmount"`     // 扣除额(单位为元)
    TotalPriceIncludingTax string `json:"totalPriceIncludingTax"` // 含税总金额(单位为元)
    TotalTax           string `json:"totalTax"`            // 税额(单位为元)
    TotalPrice         string `json:"totalPrice"`          // 不含税总金额(单位为元)
    GoodsDetail        string `json:"goodsDetail"`         // 商品明细(1-1024字符)
    NotifyMobileNo     string `json:"notifyMobileNo"`      // 消费者手机号(1-16字符)
    NotifyEmail        string `json:"notifyEmail"`         // 消费者邮箱(1-64字符)
    NotifyUrl          string `json:"notifyUrl"`           // 商户通知地址(1-128字符)
    NotifyMerEmail     string `json:"notifyMerEmail"`      // 商户邮箱(1-64字符)
    PdfUrl             string `json:"pdfUrl"`              // PDF下载链接(1-128字符)
    PdfPreviewUrl      string `json:"pdfPreviewUrl"`       // PDF预览链接(1-128字符)
    ErrMsg             string `json:"errMsg"`              // 撤销结果信息(1-256字符)
    SubList            string `json:"subList"`             // 子订单信息(JsonArray格式的字符串)
}

// BatchQueryInvoiceRequest 批量查询发票请求参数
type BatchQueryInvoiceRequest struct {
	BaseRequest
	MerchantID         string `json:"merchantId"`         // 银商商户号
	TerminalID         string `json:"terminalId"`         // 银商终端号
	StartDate          string `json:"startDate"`          // 开始日期（yyyyMMdd）
	EndDate            string `json:"endDate"`            // 结束日期（yyyyMMdd）
	PageNum            int    `json:"pageNum"`            // 页码（从1开始）
	PageSize           int    `json:"pageSize"`           // 每页条数（最大100）
	Status             string `json:"status,omitempty"`   // 发票状态（可选）
}

// BatchQueryInvoiceResponse 批量查询发票响应参数
type BatchQueryInvoiceResponse struct {
	BaseResponse
	TotalCount         int                `json:"totalCount"`         // 总记录数
	TotalPage          int                `json:"totalPage"`          // 总页数
	CurrentPage        int                `json:"currentPage"`        // 当前页码
	InvoiceList        []InvoiceListItem  `json:"invoiceList"`        // 发票列表
}

// InvoiceListItem 批量查询发票列表项
type InvoiceListItem struct {
	MerOrderID             string  `json:"merOrderId"`             // 商户订单号
	InvoiceNo              string  `json:"invoiceNo"`              // 发票号码
	InvoiceCode            string  `json:"invoiceCode"`            // 发票代码
	InvoiceDate            string  `json:"invoiceDate"`            // 开票日期
	Status                 string  `json:"status"`                 // 发票状态
	TotalPriceIncludingTax float64 `json:"totalPriceIncludingTax"` // 含税金额（元）
	BuyerName              string  `json:"buyerName"`              // 购买方名称
}