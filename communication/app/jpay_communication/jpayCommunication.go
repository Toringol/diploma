package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/Toringol/diploma/tree/master/communication/app/model"
	"github.com/Toringol/diploma/tree/master/communication/config"
	"github.com/spf13/viper"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	conn, err := net.Dial("tcp", viper.GetString("terminalIP"))
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	xmlFile, err := os.Open("../../config/transactions.yml")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	var transactions model.Transactions

	err = xml.Unmarshal(byteValue, &transactions)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	for i, transaction := range transactions.Items {
		curTransactionRequest, err := xml.Marshal(transaction)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}

		if n, err := conn.Write(curTransactionRequest); n == 0 || err != nil {
			log.Fatalf("%s", err.Error())
		}

		buff := make([]byte, (1024 * 4))
		n, err := conn.Read(buff)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}

		var curTransactionResponse model.TransactionResponse

		err = xml.Unmarshal(buff[0:n], &curTransactionResponse)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}

		if curTransactionResponse.XMLName.Local == "transaction" &&
			curTransactionResponse.Receipt.Currency == transaction.Currency &&
			curTransactionResponse.Receipt.Tstatus == "approved" &&
			curTransactionResponse.Receipt.RespCode == "000" {
			fmt.Println("Success transaction - ", i)
		} else {
			fmt.Println("Failed transaction - ", i)
			fmt.Println(curTransactionResponse.ErrorStack)
		}

	}

}
