package model

// GenerateQRCodeRequest 生成开票二维码请求参数
type GenerateQRCodeRequest struct {
	BaseRequest
    InvoiceMaterial   string `json:"invoiceMaterial"`   // 发票材质: 纸质发票：PAPER 电子发票：ELECTRONIC
    InvoiceType       string `json:"invoiceType"`       // 发票类型: 普通发票：PLAIN 增值税专用发票：VAT
    MerchantId        string `json:"merchantId"`        // 银商商户号(15位)
    TerminalId        string `json:"terminalId"`        // 银商终端号(8位)
    MerOrderDate      string `json:"merOrderDate"`      // 开票订单日期(格式yyyyMMdd)
    MerOrderId        string `json:"merOrderId"`        // 开票订单号(1-32位)
    Amount            int64  `json:"amount"`            // 开票金额(单位为分)
    DeductionAmount   int64  `json:"deductionAmount"`   // 扣除额(单位为分)
    GoodsDetail       string `json:"goodsDetail"`       // 商品明细(JSON字符串，最大15000字符)
    Remark            string `json:"remark,omitempty"`            // 备注(1-128字符)
    ExpireTime        int    `json:"expireTime"`        // 过期时间(分钟)
    NotifyUrl         string `json:"notifyUrl,omitempty"`         // 商户通知地址(1-256字符)
    ReceiveMobileNo   string `json:"receiveMobileNo,omitempty"`   // 消费者接收二维码手机号(1-16字符)
    ReceiveEmail      string `json:"receiveEmail,omitempty"`      // 消费者接收二维码邮箱(1-64字符)
    StoreId           string `json:"storeId,omitempty"`           // 门店号(1-32字符)
    StoreName         string `json:"storeName,omitempty"`         // 门店名称(1-128字符)
    Payee             string `json:"payee,omitempty"`             // 收款人(1-8字符)
    Checker           string `json:"checker,omitempty"`           // 复核人(1-8字符)
    Drawer            string `json:"drawer,omitempty"`            // 开票人(1-8字符)
    OperateId         string `json:"operateId,omitempty"`         // 营运ID(1-18字符)
    VehicleNo         string `json:"vehicleNo,omitempty"`         // 车牌号(1-10字符)
    OutTradeNo        string `json:"outTradeNo,omitempty"`        // 三方自定义的订单号(1-64字符)
    NotFaceRemark     string `json:"notFaceRemark,omitempty"`     // 非票面备注(1-256字符)
    SpecialInvoiceFlag string `json:"specialInvoiceFlag,omitempty"` // 特殊票种标识(2字符)
    InvoiceSpecialInfo string `json:"invoiceSpecialInfo,omitempty"` // 特殊票种信息(JSON字符串，最大15000字符)
    ConfirmIssue      string `json:"confirmIssue,omitempty"`      // 全电成品油单价过低是否确认开具标识(2字符)
    InvoiceBalanceInfo string `json:"invoiceBalanceInfo,omitempty"` // 数电发票差额扣除额凭证明细(JSON数组)
}


// GenerateQRCodeResponse 生成开票二维码响应参数
type GenerateQRCodeResponse struct {
	BaseResponse
    QRCodeID        string `json:"qrCodeId"`         // 二维码编号(1-40字符)
    QRCode          string `json:"qrCode"`           // 二维码(1-256字符)
    ShortQRCode     string `json:"shotQrCode"`       // 二维码短地址(1-256字符)
    AuditQRCode     string `json:"auditQrCode"`      // 审核开票二维码(1-256字符)
    AuditShortQRCode string `json:"auditShortQrCode"` // 审核开票二维码短地址(1-256字符)
    ExpireDate      string `json:"expireDate"`       // 二维码过期时间(1-32字符)
    Status          string `json:"status"`           // 二维码状态: PENDING:待开具 CLOSED:已关闭 CANCELED：已撤销 ISSUING:开具中 ISSUED:已开具 REVERSING:红冲中 REVERSED：已红冲 SPLITED:已拆分
    SubList         []any `json:"subList"`          // 子订单信息(JsonArray格式的字符串)
}


// PreQRCodeConfigRequest 预制二维码配置请求参数
type PreQRCodeConfigRequest struct {
	BaseRequest
	MerchantID         string `json:"merchantId"`         // 银商商户号
	TerminalID         string `json:"terminalId"`         // 银商终端号
	QRCodeID           string `json:"qrCodeId"`           // 预制二维码ID
	InvoiceMaterial    string `json:"invoiceMaterial"`    // 发票材质
	InvoiceType        string `json:"invoiceType"`        // 发票类型
	Amount             int64  `json:"amount,omitempty"`   // 预设金额（分）
	Status             string `json:"status"`             // 状态（ENABLE-启用/DISABLE-禁用）
}

// PreQRCodeConfigResponse 预制二维码配置响应参数
type PreQRCodeConfigResponse struct {
	BaseResponse
	QRCodeID         string `json:"qrCodeId"`          // 二维码ID
	Status           string `json:"status"`            // 配置后状态
	UpdateTime       string `json:"updateTime"`        // 更新时间
}