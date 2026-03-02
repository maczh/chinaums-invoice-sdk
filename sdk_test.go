package chinaumsinvoicesdk

import (
	"testing"

	"github.com/maczh/chinaums-invoice-sdk/api"
	"github.com/maczh/chinaums-invoice-sdk/client"
	"github.com/maczh/chinaums-invoice-sdk/model"
	"github.com/maczh/chinaums-invoice-sdk/util"
)

func TestQueryStatus(t *testing.T) {
	client.NewClient("test")
	req := model.QueryInvoiceStatusRequest{
		MerOrderDate: "20260227",
		MerOrderID:   "JHP20260227152305001",
	}
	resp, err := api.QueryInvoiceStatus(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(util.ToJsonString(resp))
}

func TestAsyncIssue(t *testing.T) {
	client.NewClient("test")
	req := model.AsyncIssueRequest{
		InvoiceMaterial: api.InvoiceMaterialElectronic, // 电子发票
		InvoiceType:     api.InvoiceTypePlain,          // 专用发票
		MerOrderDate:    "20260227",
		MerOrderId:      "JHP20260227152305001",
		BuyerTaxCode:    "914419006681800001",
		BuyerTelephone:  "13950280566",
		NotifyEmail:     "maczh@qq.com",
		BuyerName:       "福州寄海测试有限公司",
		BuyerAddress:    "福建省福州市鼓楼区",
		Amount:          13850,
		BuyerBank:       "招商银行福州广达支行",
		BuyerAccount:    "6212264100011335373",
		GoodsDetail:     "[{\"index\":\"1\",\"attribute\":\"0\",\"discountIndex\":\"\",\"name\":\"餐 饮 服 务\",\"sn\":\"3070401000000000000\",\"taxRate\":\"6\",\"priceIncludingTax\":\"138.50\",\"quantity\":\"1\",\"unit\":\"\",\"model\":\"\"}]",
	}
	resp, err := api.AsyncIssue(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(util.ToJsonString(resp))
}

func TestQRCode(t *testing.T) {
	client.NewClient("test")
	req := model.GenerateQRCodeRequest{
		InvoiceMaterial: api.InvoiceMaterialElectronic, // 电子发票
		InvoiceType:     api.InvoiceTypePlain,          // 专用发票
		MerOrderDate:    "20260227",
		MerOrderId:      "JHP20260227152305002",
		NotifyUrl:       "https://baidu.com",
		Amount:          13850,
		GoodsDetail:     "[{\"index\":\"1\",\"attribute\":\"0\",\"discountIndex\":\"\",\"name\":\"餐 饮 服 务\",\"sn\":\"3070401000000000000\",\"taxRate\":\"6\",\"priceIncludingTax\":\"138.50\",\"quantity\":\"1\",\"unit\":\"\",\"model\":\"\"}]",
	}
	resp, err := api.GenerateQRCode(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(util.ToJsonString(resp))
}

func TestFuzzyQueryInvoiceTitle(t *testing.T) {
	client.NewClient("test")
	req := model.FuzzyQueryInvoiceTitleRequest{
		Name: "寄海",
	}
	resp, err := api.FuzzyQueryInvoiceTitle(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(util.ToJsonString(resp))
}

func TestRedReverseInvoice(t *testing.T) {
	client.NewClient("test")
	req := model.RedReverseInvoiceRequest{
		MerOrderDate: "20260302",
		MerOrderID:   "3eb07413-416c-4be5-bca9-c9fe29d7663a-776dzq",
	}
	resp, err := api.RedReverseInvoice(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(util.ToJsonString(resp))
}
