package client

import "time"

var configs = map[string]Config{
	"prod": {
		BaseURL: "https://api.chinaums.com",
		MsgSrc: "",
		SignKey: "",
		MerchantID: "",
		TerminalID: "",
		Timeout:    30 * time.Second,
	},
	"test": {
		BaseURL:    "https://mobl-test.chinaums.com/fapiao-api-test/",
		MsgSrc:     "",
		SignKey:    "",
		MerchantID: "",
		TerminalID: "",
		Timeout:    30 * time.Second,
	},
}

