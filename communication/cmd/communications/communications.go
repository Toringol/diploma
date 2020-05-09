package main

import (
	"log"

	"github.com/Toringol/diploma/tree/master/communication/app/uartCommunication"

	"github.com/Toringol/diploma/tree/master/communication/app/jpayCommunication"

	"github.com/Toringol/diploma/tree/master/communication/config"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	conn, err := jpayCommunication.TerminalConnect()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	defer conn.Close()

	port, err := uartCommunication.ArmConnect()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	transactions, err := jpayCommunication.GetConfigTransactions()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	err = jpayCommunication.TestingTerminalProcess(conn, port, transactions)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

}
