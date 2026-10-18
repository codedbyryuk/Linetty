package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleClients(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected: ", conn.RemoteAddr())

	_, err := conn.Write([]byte("Welcome to Linetty\n"))
	if err != nil {
		fmt.Println("Error sending message: ", err)
		return
	}

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading message: ", err)
			return
		}

		fmt.Printf("Client says: %s", msg)

		reply := "server recieved: " + msg

		_, err = conn.Write([]byte(reply))

		if err != nil {
			fmt.Println("Error sending reply: ", err)
			return
		}

		fmt.Println("Reply sent!")
	}

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error creating connection: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening to port 8080")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err)
			continue
		}

		go handleClients(conn)
	}

}
