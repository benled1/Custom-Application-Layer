package http

import (
	"custom-webserver/http/utils"
	"fmt"
	"net"
)

type HTTPServer struct {
}

func (s *HTTPServer) HandleRequest(conn net.Conn) error {
	newHttpRequest, err := utils.NewHTTPRequest(conn)
	if err != nil {
		return fmt.Errorf("error while creating a new HTTPRequest: %w", err)
	}
	fmt.Println("Made a new request of method = ", newHttpRequest.Method)
	return nil
	// var msg string = "HTTP/1.1 200 OK\nContent-Length: 13\n\nHello, world!"
	// _, writeErr := conn.Write([]byte(msg))
	// if writeErr != nil {
	// 	fmt.Println("Error while writing to connection...")
	// 	return nil
	// }
}

func (s *HTTPServer) addHeaders() {

}
