package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"os"
)

func main() {
	// Create a connection with remote machine
	conn, connErr := net.Dial("tcp", "[::1]:8081")
	if connErr != nil {
		// log.Fatal(connErr)
		fmt.Printf("Could not connect to %s, target machine may be down or refused connection", conn.RemoteAddr())
		conn.Close()
		os.Exit(0)
	}
	defer conn.Close()

	fmt.Printf("\nConnected.\n")
	fmt.Println("Connection Details:")
	fmt.Printf("Remote Address: %s\n", conn.RemoteAddr())
	fmt.Printf("Remote Address: %s\n", conn.LocalAddr())

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)

	for {
		// read one line (ended with \n or \r\n)
		line, _ := tp.ReadLine()
		fmt.Println(line)
		if line == "" {
			break
		}
	}
}
