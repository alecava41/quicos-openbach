package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	for {
		var number int
		_, err := fmt.Fscanf(conn, "%d\n", &number)
		if err != nil {
			fmt.Println("Error receiving data from client:", err)
			return
		}

		fmt.Printf("Received: %d\n", number)

		// Send the same number back to the client
		_, err = fmt.Fprintf(conn, "%d\n", number)
		if err != nil {
			fmt.Println("Error sending data to client:", err)
			return
		}

		fmt.Printf("Sent: %d\n", number)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	fmt.Println("Server started, waiting for connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
