package model

import "encoding/xml"

type Transactions struct {
	XMLName xml.Name             `xml:"transactions"`
	Items   []TransactionRequest `xml:"transaction"`
}
