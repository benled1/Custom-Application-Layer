package utils

import (
	"fmt"
	"net"
	"strings"
)

type HTTPRequest struct {
	URL      string
	Method   string
	Protocol string
	Header   string
	Body     string
}

func NewHTTPRequest(conn net.Conn) (*HTTPRequest, error) {
	newRequest := new(HTTPRequest)
	requestMsg, err := newRequest.readRequest(conn)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request connection: %w", err)
	}
	// here we want to parse the requestMsg and then fill the URL, Method, Header, and Body
	fmt.Println(requestMsg)
	newRequest.parseRequest(requestMsg)
	fmt.Println("r.Method = ", newRequest.Method)
	fmt.Println("r.URL = ", newRequest.URL)
	fmt.Println("r.Protocol = ", newRequest.Protocol)
	fmt.Println("r.Header = ", newRequest.Header)
	fmt.Println("r.Body = ", newRequest.Body)
	fmt.Println("End of request...")
	return newRequest, nil
}

func (r *HTTPRequest) readRequest(conn net.Conn) (string, error) {
	var buff []byte = make([]byte, 1024)
	n, readErr := conn.Read(buff)
	if readErr != nil {
		return "", fmt.Errorf("error while reading from connection: %w", readErr)
	}
	var recvedData string = string(buff[:n])
	return recvedData, nil
}

func (r *HTTPRequest) parseRequest(requestMsg string) error {
	var requestSlice []string = strings.Split(requestMsg, "\r\n")
	var line0 []string = strings.Split(requestSlice[0], " ")
	r.Method = line0[0]
	r.URL = line0[1]
	r.Protocol = line0[2]

	var bodyStartIndex int = -1
	for i := 0; i < len(requestSlice); i++ {
		if requestSlice[i] == "" {
			bodyStartIndex = i + 1
		}
	}
	if bodyStartIndex != -1 && bodyStartIndex < len(requestSlice) {
		r.Body = strings.Join(requestSlice[bodyStartIndex:], "\r\n")
		r.Header = strings.Join(requestSlice[1:bodyStartIndex-1], "\r\n")
	} else {
		r.Header = strings.Join(requestSlice[1:], "\r\n")
		r.Body = ""
	}
	return nil
}
