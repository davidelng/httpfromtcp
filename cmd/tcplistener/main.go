package main

import (
	"fmt"
	"log"
	"net"

	"github.com/davidelng/httpfromtcp/internal/request"
)

const port = ":42069"

func main() {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}
	defer ln.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())

		req, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatalf("error parsing request: %s\n", err.Error())
		}

		fmt.Println("Request line:")
		fmt.Printf("- Method: %s\n", req.RequestLine.Method)
		fmt.Printf("- Target: %s\n", req.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n", req.RequestLine.HttpVersion)
		if len(req.Headers) > 0 {
			fmt.Println("Headers:")
		}
		for key, value := range req.Headers {
			fmt.Printf("- %s: %s\n", key, value)
		}
		if len(req.Body) > 0 {
			fmt.Println("Body:")
			fmt.Printf("%s", req.Body)
		}
	}
}
