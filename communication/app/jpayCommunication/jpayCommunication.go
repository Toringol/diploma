package jpayCommunication

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/Toringol/diploma/tree/master/communication/app/model"
	"github.com/Toringol/diploma/tree/master/communication/app/uartCommunication"
	"github.com/schleibinger/sio"
	"github.com/spf13/viper"
)

func TerminalConnect() (net.Conn, error) {
	conn, err := net.Dial("tcp", viper.GetString("terminalIP")+viper.GetString("terminalPort"))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetConfigTransactions() (model.Transactions, error) {
	xmlFile, err := os.Open(viper.GetString("transactionsConfigPath"))
	if err != nil {
		return model.Transactions{}, err
	}

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return model.Transactions{}, err
	}

	var transactions model.Transactions

	err = xml.Unmarshal(byteValue, &transactions)
	if err != nil {
		return model.Transactions{}, err
	}

	return transactions, nil
}

func TestingTerminalProcess(conn net.Conn, port *sio.Port, transactions model.Transactions) error {

	for i, transaction := range transactions.Items {
		curTransactionRequest, err := xml.Marshal(transaction)
		if err != nil {
			return err
		}

		if n, err := conn.Write(curTransactionRequest); n == 0 || err != nil {
			return err
		}

		err = uartCommunication.ArmMove(port)
		if err != nil {
			return err
		}

		ch := make(chan []byte)
		eCh := make(chan error)

		go func(ch chan []byte, eCh chan error) {
			for {
				data := make([]byte, (1024 * 4))
				n, err := conn.Read(data)
				if err != nil {
					eCh <- err
				}

				var checkTransaction model.TransactionResponse
				if err := xml.Unmarshal(data[0:n], &checkTransaction); err != nil ||
					checkTransaction.XMLName.Local != "transaction" {
					continue
				}

				ch <- data[0:n]
			}
		}(ch, eCh)

		for {
			select {
			case data := <-ch:
				var curTransactionResponse model.TransactionResponse

				err = xml.Unmarshal(data, &curTransactionResponse)
				if err != nil {
					return err
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
				break
			case err := <-eCh:
				return err
			}
		}
	}

	return nil
}
