package main

import (
	"custom-application-layer/http"
	"custom-application-layer/tcp"
)

func main() {
	var server tcp.TCPServer = tcp.TCPServer{ApplicationServer: http.HTTPServer{}}
	server.Start()
}
