package main

import "fmt"

const (
	connHost = "*"
	connPort = "61745"
	connType = "tcp"
)

func main() {
	fmt.Println("Arrancando P!TAS Server (" + connType +") on " + connHost + ":" + connPort)
	
}