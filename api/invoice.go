package api

import (
	"github.com/maczh/chinaums-invoice-sdk/client"
	"github.com/maczh/chinaums-invoice-sdk/model"
)

// DirectIssue 直接开票（文档5.2）
func DirectIssue(req model.DirectIssueRequest) (*model.DirectIssueResponse, error) {
	req.SetMsgType(MSG_TYPE_ISSUE)
	resp := &model.DirectIssueResponse{}
	if err := client.SendRequest[model.DirectIssueRequest, model.DirectIssueResponse](req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// AsyncIssue 异步开票（文档6）
func AsyncIssue(req model.AsyncIssueRequest) (*model.AsyncIssueResponse, error) {
	req.SetMsgType(MSG_TYPE_ASYNC_ISSUE)
	resp := &model.AsyncIssueResponse{}
	if err := client.SendRequest[model.AsyncIssueRequest, model.AsyncIssueResponse](req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
