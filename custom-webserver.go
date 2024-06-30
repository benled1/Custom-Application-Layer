package main

import (
	"custom-webserver/http"
	"custom-webserver/tcp"
)

func main() {
	httpServer := http.HTTPServer{}
	var server tcp.TCPServer = tcp.TCPServer{ApplicationServer: httpServer}
	server.Start()
}
