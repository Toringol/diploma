package model

type TransactionResponse struct {
	Status  string `xml:"status"`
	Receipt Receipt
}

type Receipt struct {
	Tid      string `xml:"tid"`
	Mid      string `xml:"mid"`
	State    string `xml:"state"`
	Tstatus  string `xml:"tstatus"`
	Currency string `xml:"currency"`
	RespCode string `xml:"resp-code"`
}
