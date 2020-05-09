package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/schleibinger/sio"
)

func main() {
	// устанавливаем соединение
	port, err := sio.Open("/dev/ttyACM1", syscall.B9600)
	if err != nil {
		log.Fatal(err)
	}

	input := ""

	for {
		fmt.Scanln(&input)

		if input == "data" {
			// отправляем данные
			_, err = port.Write([]byte("1"))
			if err != nil {
				log.Fatal(err)
			}
		}
		if input == "exit" {
			break
		}
	}
}
