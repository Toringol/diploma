package uartCommunication

import (
	"syscall"

	"github.com/schleibinger/sio"
)

func ArmConnect() (*sio.Port, error) {
	port, err := sio.Open("/dev/ttyACM1", syscall.B9600)
	if err != nil {
		return nil, err
	}

	return port, nil
}

func ArmMove(port *sio.Port) error {
	_, err := port.Write([]byte("someData"))
	if err != nil {
		return err
	}

	return nil
}
