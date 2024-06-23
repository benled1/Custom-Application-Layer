package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	var msg string = "HTTP/1.1 200 OK\nContent-Length: 13\n\nHello, world!"
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
