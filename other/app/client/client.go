package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: client <server-address> <server-port>")
		return
	}

	serverAddr := os.Args[1]
	serverPort, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid server port:", err)
		return
	}

	server := fmt.Sprintf("%s:%d", serverAddr, serverPort)

	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	rand.Seed(time.Now().UnixNano())

	for {
		number := rand.Intn(100) // Generate a random number between 0 and 99

		_, err := fmt.Fprintf(conn, "%d\n", number)
		if err != nil {
			fmt.Println("Error sending data to server:", err)
			return
		}

		fmt.Printf("Sent: %d\n", number)

		// Wait for server response
		var response int
		_, err = fmt.Fscanf(conn, "%d\n", &response)
		if err != nil {
			fmt.Println("Error receiving data from server:", err)
			return
		}

		fmt.Printf("Received: %d\n", response)

		time.Sleep(1 * time.Second) // Wait for 1 second before sending the next number
	}
}
