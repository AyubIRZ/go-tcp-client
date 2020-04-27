package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

const TCPServer = "localhost:6060"

func main() {
	wg := &sync.WaitGroup{}

	fmt.Println("* * * TCP client started * * *")
	fmt.Println("==============================")

	conn := initiateTCPConn(TCPServer)

	wg.Add(2)
	go receiveMessage(conn, wg)
	go sendMessage(conn, wg)

	wg.Wait()
}


// initiateTCPConn makes a TCP connection session to the server and returns the conn instance.
func initiateTCPConn(TCPServer string) net.Conn {
	tcpAddr, err := net.ResolveTCPAddr("tcp", TCPServer)
	if err != nil {
		log.Fatal("TCP resolve failed: ", err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("TCP connection init failed: ", err)
	}

	return conn
}


// receiveMessage is a worker reading from TCP socket and writing the message to stdout.
func receiveMessage(conn net.Conn, wg *sync.WaitGroup) {
	buf := bufio.NewReader(conn)

	for {
		msg, err := buf.ReadString('\n')
		if err != nil {
			log.Fatal("TCP connection read failed: ", err)
		}
		print("\033[F\033[K\033[F\r\033[K")
		fmt.Print("<Server>: ", msg[:len(msg) - 1])
		fmt.Print("\n---------------------------\nEnter your message: ")
	}

	conn.Close()
	wg.Done()
}


// sendMessage is a worker reading from stdin and writing the message to the TCP socket.
func sendMessage(conn net.Conn, wg *sync.WaitGroup) {
	for {
		fmt.Print("\n---------------------------\nEnter your message: ")

		msg, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal("Reading from chat input failed: ", err)
		}

		if msg == "\r\n"{
			print("\r\033[K\033[F\r\033[K\033[F\r\033[K\033[F")
			continue
		}

		fmt.Print("\r\033[K\033[F\r\033[K\033[F\r\033[K")
		fmt.Print("<YOU>: ", msg)
		_, _ = fmt.Fprint(conn, msg)
	}

	wg.Done()
}

