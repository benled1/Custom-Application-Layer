package main

import (
	"fmt"
	"net"
)

// take a look at this article: https://okanexe.medium.com/the-complete-guide-to-tcp-ip-connections-in-golang-1216dae27b5a

func handleConnection(conn net.Conn) {
	var msg string = "hello from the server"
	_, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Error while writing to connection...")
		return
	}
}

func main() {
	fmt.Println("Server listening on port 8080...")
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error on listener")
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in message from client")
			continue
		}
		go handleConnection(conn)
	}

}
