package model

import "encoding/xml"

type TransactionResponse struct {
	XMLName    xml.Name   `xml:"transaction"`
	Status     string     `xml:"status"`
	Receipt    Receipt    `xml:"receipt"`
	ErrorStack ErrorStack `xml:"error-stack"`
}

type Receipt struct {
	Tid      string `xml:"tid"`
	Mid      string `xml:"mid"`
	State    string `xml:"state"`
	Tstatus  string `xml:"tstatus"`
	Currency string `xml:"currency"`
	RespCode string `xml:"resp-code"`
}

type ErrorStack struct {
	XMLName xml.Name `xml:"error-stack"`
	Error   string   `xml:"error"`
}
