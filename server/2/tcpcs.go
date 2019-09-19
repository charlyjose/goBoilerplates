package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// Client ...
type Client struct {
	socket net.Conn    // Connection details
	data   chan []byte // Data from a client
}

// ClientManager ...
type ClientManager struct {
	clients    map[*Client]bool // Clients details
	broadcast  chan []byte      // broadcast message
	register   chan *Client     // Register client
	unregister chan *Client     // Unregister client
}

func (manager *ClientManager) start() {
	for {
		select {
		case client := <-manager.register:
			manager.clients[client] = true
			fmt.Printf("Added client: %s\n", client.socket.RemoteAddr())
		case client := <-manager.unregister:
			if _, ok := <-manager.unregister; !ok {
				close(client.data)
				fmt.Printf("[%s] Connection teminated\n", client.socket.RemoteAddr())
				delete(manager.clients, client)
			}
		case message := <-manager.broadcast:
			for client := range manager.clients {
				select {
				case client.data <- message:
				default:
					close(client.data)
					delete(manager.clients, client)
				}
			}
		}
	}
}

func (client *Client) receive() {
	for {
		message := make([]byte, 4096)
		length, rErr := client.socket.Read(message)
		if rErr != nil {
			client.socket.Close()
		}
		if length > 0 {
			fmt.Printf("(%s) Received: %s", client.socket.RemoteAddr(), message)
		}
	}
}

func (manager *ClientManager) send(client *Client) {
	defer client.socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			client.socket.Write(message)
		}
	}
}

func startServerMode(port int) {
	listener, lisErr := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))
	if lisErr != nil {
		log.Fatal(lisErr)
		listener.Close()
		os.Exit(0)
	}
	defer listener.Close()

	fmt.Printf("Server started listening at %s\n", listener.Addr())

	manager := ClientManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	// Communication manager routine
	go manager.start()

	// Listener
	for {
		// Look for new connections
		conn, connErr := listener.Accept()
		if connErr != nil {
			log.Fatal(connErr)
			conn.Close()
		}

		client := &Client{socket: conn, data: make(chan []byte)}
		// setting register
		manager.register <- client
		// send and receive
		go client.receive()     // receive message routine
		go manager.send(client) // send message routine
	}
}

func startClientMode(port int) {
	fmt.Printf("Connecting...\n")
	conn, conErr := net.Dial("tcp", "0.0.0.0:"+strconv.Itoa(port))
	if conErr != nil {
		log.Fatal(conErr)
		conn.Close()
		os.Exit(0)
	}
	defer conn.Close()

	fmt.Printf("Connected\n")

	client := &Client{socket: conn}
	go client.receive()
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(strings.TrimRight(message, "\n")))
	}
}

func main() {
	mode := flag.String("mode", "server", "Select a mode: server/client")
	port := flag.Int("port", 8081, "Select a port")

	flag.Parse()

	if strings.ToLower(*mode) == "server" {
		startServerMode(*port)
	} else {
		startClientMode(*port)
	}
}
