package api

const (
	// 消息类型
	MSG_TYPE_ISSUE = "issue" // 开具发票
	MSG_TYPE_ASYNC_ISSUE = "async.issue" // 异步开具发票
	MSG_TYPE_QRCODE = "get.qrcode" // 获取开票二维码
	MSG_TYPE_QUERY_STATUS = "query" // 查询发票状态
	MSG_TYPE_REVERSE = "reverse" // 红冲发票
	MSG_TYPE_INVALID = "invalid" // 作废发票
	MSG_TYPE_FUZZY_QUERY = "query.fuzzy.title" // 抬头模糊查询

	// 发票材质
	InvoiceMaterialPaper      = "PAPER"      // 纸质发票
	InvoiceMaterialElectronic = "ELECTRONIC" // 电子发票

	// 发票类型
	InvoiceTypePlain = "PLAIN" // 普通发票
	InvoiceTypeVAT   = "VAT"   // 增值税发票

	// 发票状态
	ISSUE_STATUS_ISSUED = "ISSUED" // 已开具
	ISSUE_STATUS_ISSUING = "ISSUING" // 开具中
	ISSUE_STATUS_CLOSED = "CLOSED" // 已关闭
	ISSUE_STATUS_SPLITED = "SPLITED" // 已拆分
	ISSUE_STATUS_CANCELLED = "CANCELLED" // 已取消
	ISSUE_STATUS_REVERSED = "REVERSED" // 已红冲
	ISSUE_STATUS_REVERSING = "REVERSING" // 红冲中
	ISSUE_STATUS_INVALIDED = "INVALIDED" // 已作废
	ISSUE_STATUS_INVALIDING = "INVALIDING" // 作废中
)
