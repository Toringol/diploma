package model

import "encoding/xml"

type TransactionRequest struct {
	XMLName  xml.Name `xml:"transaction"`
	Type     string   `xml:"type"`
	Currency string   `xml:"currency"`
	Amount   string   `xml:"amount"`
}
