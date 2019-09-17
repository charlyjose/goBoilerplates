package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func server(port int) {
	listener, LisErr := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))
	if LisErr != nil {
		log.Fatal(LisErr)
		listener.Close()
	}

	fmt.Printf("Server started listening at %s", listener.Addr())

}

func client(port int) {
	
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
