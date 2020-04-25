package main

import (
	"fmt"
	"log"
	"net"
)

const TCPServer = "localhost:6060"

func main() {
	fmt.Println("* * * TCP client started * * *")

	tcpAddr, err := net.ResolveTCPAddr("tcp", TCPServer)
	if err != nil {
		log.Fatal("TCP resolve failed: ", err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("TCP conn failed: ", err)
	}
	defer conn.Close()

	fmt.Fprint(conn, "Hello from TCP client!!!\n")
}