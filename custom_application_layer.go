package main

import (
	"custom-application-layer/tcp"
)

func main() {
	var server tcp.TCPServer = tcp.TCPServer{ApplicationServer: "http"}
	server.Start()
}
