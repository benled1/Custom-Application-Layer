package http

import (
	"fmt"
	"net"
)

func HandleRequest(conn net.Conn) {
	var buff []byte = make([]byte, 1024)
	n, readErr := conn.Read(buff)
	if readErr != nil {
		fmt.Println("Error while reading from connection ", readErr)
		return
	}
	var recvedData string = string(buff[:n])
	fmt.Println(recvedData)

	var msg string = "HTTP/1.1 200 OK\nContent-Length: 13\n\nHello, world!"
	_, writeErr := conn.Write([]byte(msg))
	if writeErr != nil {
		fmt.Println("Error while writing to connection...")
		return
	}
}
