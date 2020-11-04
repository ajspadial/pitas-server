package main

import (
	"fmt"
	"net"
)

const (
	connHost = "*"
	connPort = "61745"
	connType = "tcp"
)

func main() {
	fmt.Println("Arrancando P!TAS Server (" + connType +") on " + connHost + ":" + connPort)
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error en escucha:", err.Error())
	}
	defer l.Close();

	for {

	}
}