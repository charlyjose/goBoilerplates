package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"os"
	"strconv"
	"strings"
)

func server(port int) {
	// Listen on all interfaces
	listener, lisErr := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))
	if lisErr != nil {
		log.Fatal(LisErr)
		listener.Close()
		os.Exit(0)
	}
	defer listener.Close()

	fmt.Println("Server started listening at ", listener.Addr())

	for {
		fmt.Println("\nReady to accept new connections")

		// Accept connection on port
		conn, conErr := listener.Accept()
		if conErr != nil {
			log.Fatal(conErr)
			conn.Close()
		}
		defer conn.Close()
		fmt.Printf("\n1 Connection accepted.\n")
		fmt.Println("Connection Details:")
		fmt.Printf("Remote Address: %s\n", conn.RemoteAddr())

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

		fmt.Printf("Request:\n\n")
		for {
			// read one line (ended with \n or \r\n)
			line, _ := tp.ReadLine()
			fmt.Println(line)
			if line == "" {
				break
			}
		}
		// conn.Close()
		fmt.Println("\nConnection Closed")
	}

}

func client(port int) {
	// Create a connection with remote machine
	conn, connErr := net.Dial("tcp", ":"+strconv.Itoa(port))
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
	fmt.Printf("Local Address: %s\n", conn.LocalAddr())

	// reader := bufio.NewReader(conn)
	// tp := textproto.NewReader(reader)

	// for {
	// 	// read one line (ended with \n or \r\n)
	// 	line, _ := tp.ReadLine()
	// 	fmt.Println(line)
	// 	if line == "" {
	// 		break
	// 	}
	// }
}

func main() {
	mode := flag.String("mode", "server", "Select a mode: server/client")
	port := flag.Int("port", 8081, "Select a port")
	flag.Parse()

	if strings.ToLower(*mode) == "server" {
		server(*port)
	} else {
		client(*port)
	}

}

/*


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
