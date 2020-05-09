package jpay_communication

import (
	"log"
	"net"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	conn, err := net.Dial("tcp", "192.168.0.100")
	if err != nil {
		log.Fatal("Connection Error")
	}

}
