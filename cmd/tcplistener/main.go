package main

import (
	"fmt"
	"httpfromtcp/internal/request"
	"log"
	"net"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())

		go func(c net.Conn) {
			line, err := request.RequestFromReader(conn)
			if err != nil {
				log.Fatalf("%v", err)
			}
			fmt.Printf("Request line:\n")
			fmt.Printf("- Method: %v\n", line.RequestLine.Method)
			fmt.Printf("- Target: %v\n", line.RequestLine.RequestTarget)
			fmt.Printf("- Version: %v\n", line.RequestLine.HttpVersion)
			fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
		}(conn)
	}
}
