package model

type TransactionRequest struct {
	Type     string `xml:"type"`
	Currency string `xml:"currency"`
	Amount   string `xml:"amount"`
}
