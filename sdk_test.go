package chinaumsinvoicesdk

import (
	"testing"

	"github.com/maczh/chinaums-invoice-sdk/api"
	"github.com/maczh/chinaums-invoice-sdk/client"
	"github.com/maczh/chinaums-invoice-sdk/model"
	"github.com/maczh/chinaums-invoice-sdk/util"
)

func TestIssue(t *testing.T) {
	client.NewClient("test")
	req := model.DirectIssueRequest{
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
	resp, err := api.DirectIssue(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(util.ToJsonString(resp))
}
