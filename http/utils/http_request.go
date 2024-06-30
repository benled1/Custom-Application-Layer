package utils

import (
	"fmt"
	"net"
)

type HTTPRequest struct {
	URL    string
	Method string
	Header string
	Body   string
}

func NewHTTPRequest(conn net.Conn) (*HTTPRequest, error) {
	newRequest := new(HTTPRequest)
	requestMsg, err := newRequest.readRequest(conn)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request connection: %w", err)
	}
	// here we want to parse the requestMsg and then fill the URL, Method, Header, and Body
	fmt.Println(requestMsg)
	return newRequest, nil
}

func (r HTTPRequest) readRequest(conn net.Conn) (string, error) {
	var buff []byte = make([]byte, 1024)
	n, readErr := conn.Read(buff)
	if readErr != nil {
		return "", fmt.Errorf("error while reading from connection: %w", readErr)
	}
	var recvedData string = string(buff[:n])
	return recvedData, nil
}

func (r HTTPRequest) parseRequest(requestMsg string) error {

	return nil
}
