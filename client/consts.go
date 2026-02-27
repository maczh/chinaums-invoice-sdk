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
		MsgSrc:     "GUILIN_SWJ",
		SignKey:    "e43cb05eafba4c74821d736d161812b1",
		MerchantID: "898000000000122",
		TerminalID: "00000102",
		Timeout:    30 * time.Second,
	},
}

