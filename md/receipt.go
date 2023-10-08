package md

type AppStoreVerifyReceipt struct {
	ReceiptData            string `json:"receipt-data"`
	Password               string `json:"password"`
	ExcludeOldTransactions bool   `json:"exclude-old-transactions"`
}

type AppStoreVerifyReceiptResp struct {
	Environment string `json:"environment"`
	IsRetryable bool   `json:"is-retryable"`
	Status      int    `json:"status"`
}
