package api

import (
	"github.com/maczh/chinaums-invoice-sdk/client"
	"github.com/maczh/chinaums-invoice-sdk/model"
)

// InvoiceManageAPI 发票管理接口

// QueryInvoiceStatus 查询发票状态（文档9.1）
func QueryInvoiceStatus(req model.QueryInvoiceStatusRequest) (*model.QueryInvoiceStatusResponse, error) {
	req.SetMsgType(MSG_TYPE_QUERY_STATUS)
	resp := &model.QueryInvoiceStatusResponse{}
	if err := client.SendRequest[model.QueryInvoiceStatusRequest, model.QueryInvoiceStatusResponse](req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// RedReverseInvoice 红冲发票（文档9.5）
func RedReverseInvoice(req model.RedReverseInvoiceRequest) (*model.RedReverseInvoiceResponse, error) {
	req.SetMsgType(MSG_TYPE_REVERSE)
	resp := &model.RedReverseInvoiceResponse{}
	if err := client.SendRequest[model.RedReverseInvoiceRequest, model.RedReverseInvoiceResponse](req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// InvalidInvoice 作废发票（文档9.7）
func InvalidInvoice(req model.InvalidInvoiceRequest) (*model.InvalidInvoiceResponse, error) {
	req.SetMsgType(MSG_TYPE_INVALID)
	resp := &model.InvalidInvoiceResponse{}
	if err := client.SendRequest[model.InvalidInvoiceRequest, model.InvalidInvoiceResponse](req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// FuzzyQueryInvoiceTitle 抬头模糊查询（文档9.8）
func FuzzyQueryInvoiceTitle(req model.FuzzyQueryInvoiceTitleRequest) (*model.FuzzyQueryInvoiceTitleResponse, error){
	req.SetMsgType(MSG_TYPE_FUZZY_QUERY)
	resp := &model.FuzzyQueryInvoiceTitleResponse{}
	if err := client.SendRequest[model.FuzzyQueryInvoiceTitleRequest, model.FuzzyQueryInvoiceTitleResponse](req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}