package model

// InvoiceInspectionRequest 发票查验请求参数
type InvoiceInspectionRequest struct {
	BaseRequest
	MerchantID         string `json:"merchantId"`         // 银商商户号
	TerminalID         string `json:"terminalId"`         // 银商终端号
	InvoiceType        string `json:"invoiceType"`        // 发票类型（PLAIN/VAT/ELECTRONIC）
	InvoiceNo          string `json:"invoiceNo"`          // 发票号码
	InvoiceCode        string `json:"invoiceCode"`        // 发票代码
	InvoiceDate        string `json:"invoiceDate"`        // 开票日期（yyyyMMdd）
	CheckCode          string `json:"checkCode,omitempty"`// 校验码（后6位，普通发票必填）
	Amount             string `json:"amount,omitempty"`   // 金额（专票必填，元）
}

// InvoiceInspectionResponse 发票查验响应参数
type InvoiceInspectionResponse struct {
	BaseResponse
	InvoiceNo          string `json:"invoiceNo"`          // 发票号码
	InvoiceCode        string `json:"invoiceCode"`        // 发票代码
	InvoiceDate        string `json:"invoiceDate"`        // 开票日期
	SellerName         string `json:"sellerName"`         // 销售方名称
	SellerTaxCode      string `json:"sellerTaxCode"`      // 销售方税号
	BuyerName          string `json:"buyerName"`          // 购买方名称
	BuyerTaxCode       string `json:"buyerTaxCode"`       // 购买方税号
	TotalAmount        string `json:"totalAmount"`        // 合计金额（元）
	TotalTax           string `json:"totalTax,omitempty"` // 合计税额（元）
	Status             string `json:"status"`             // 查验状态（VALID-有效/INVALID-无效）
}

// OCRRecognitionRequest 发票OCR识别请求参数
type OCRRecognitionRequest struct {
	BaseRequest
	MerchantID         string `json:"merchantId"`         // 银商商户号
	TerminalID         string `json:"terminalId"`         // 银商终端号
	ImageURL           string `json:"imageUrl"`           // 发票图片URL
	ImageContent       string `json:"imageContent,omitempty"`// 图片Base64编码（二选一）
	InvoiceType        string `json:"invoiceType,omitempty"`// 发票类型（可选）
}

// OCRRecognitionResponse 发票OCR识别响应参数
type OCRRecognitionResponse struct {
	BaseResponse
	InvoiceNo          string `json:"invoiceNo,omitempty"` // 发票号码
	InvoiceCode        string `json:"invoiceCode,omitempty"`// 发票代码
	InvoiceDate        string `json:"invoiceDate,omitempty"`// 开票日期
	SellerName         string `json:"sellerName,omitempty"` // 销售方名称
	SellerTaxCode      string `json:"sellerTaxCode,omitempty"`// 销售方税号
	BuyerName          string `json:"buyerName,omitempty"`  // 购买方名称
	BuyerTaxCode       string `json:"buyerTaxCode,omitempty"`// 购买方税号
	TotalAmount        string `json:"totalAmount,omitempty"`// 合计金额
	RecognizeAccuracy  string `json:"recognizeAccuracy"`   // 识别准确率（%）
}