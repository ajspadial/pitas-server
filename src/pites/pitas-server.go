package main

import (
	"fmt"
	"os"
	"net"
	"log"
	"bufio"
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
		os.Exit(1)
	}
	defer l.Close();

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error de conexion:"+err.Error())
			return
		}
		fmt.Println("Cliente conectado")

		fmt.Println("Cliente conectado desde " + c.RemoteAddr().String())

		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Cliente ha salido")
		conn.Close()
		return
	}

	log.Println("Mensaje de cliente:", string(buffer[:len(buffer)-1]))

	conn.Write(buffer)

	handleConnection(conn)
}