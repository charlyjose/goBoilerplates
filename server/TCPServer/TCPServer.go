package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/textproto"
)

func main() {
	// Listen on all interfaces
	listener, LisErr := net.Listen("tcp", "[::1]:8081")
	if LisErr != nil {
		log.Fatal(LisErr)
	}

	fmt.Println("Server started listening at ", listener.Addr())

	for {
		fmt.Println("\nReady to accept new connections")

		// Accept connection on port
		conn, conErr := listener.Accept()
		if conErr != nil {
			log.Fatal(conErr)
			conn.Close()
		}
		// defer conn.Close()

		fmt.Println("\nConnection Details:")
		fmt.Println("Remote Address: ", conn.RemoteAddr(), "\n")

		/*
			for {
				message, _ := bufio.NewReader(conn).ReadString('\n')
				// output message received
				fmt.Print("Message Received:", string(message))

				// sample process for string received
				newmessage := strings.ToUpper(message)
				// send new string back to client
				conn.Write([]byte(newmessage + "\n"))
			}
		*/

		reader := bufio.NewReader(conn)
		tp := textproto.NewReader(reader)

		fmt.Println("Request: \n")
		for {
			// read one line (ended with \n or \r\n)
			line, _ := tp.ReadLine()

			fmt.Println(line)
			if line == "" {
				break
			}
		}
		conn.Close()
		fmt.Println("\nConnection Closed")
	}

}

func mainu() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	var buf bytes.Buffer
	io.Copy(&buf, conn)
	fmt.Println("total size:", buf.Len())
}

/*


GET / HTTP/1.1
Cache-Control: max-age=0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,* /*;q=0.8
Accept-Language: en-US
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18362
Accept-Encoding: gzip, deflate
Host: localhost:8081
Connection: Keep-Alive
DNT: 1


*/
