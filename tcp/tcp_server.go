package tcp

import (
	"custom-application-layer/http"
	"fmt"
	"net"
)

type TCPServer struct {
	ApplicationServer string
}

func (s TCPServer) Start() {

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

		if s.ApplicationServer == "http" {
			http.HandleRequest(conn)
		}
	}
}
