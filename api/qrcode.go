package api

import (
	"github.com/maczh/chinaums-invoice-sdk/client"
	"github.com/maczh/chinaums-invoice-sdk/model"
)

// QRCodeAPI 二维码相关接口

// GenerateQRCode 生成开票二维码（文档7.2）
func GenerateQRCode(req model.GenerateQRCodeRequest) (*model.GenerateQRCodeResponse, error) {
	req.SetMsgType(MSG_TYPE_QRCODE)
	resp := &model.GenerateQRCodeResponse{}
	if err := client.SendRequest[model.GenerateQRCodeRequest, model.GenerateQRCodeResponse](req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
