package main

import (
	"flag"
	"net"
	"fmt"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on (default 8080)")
	flag.Parse()	// initiate the server

	startServer(*port)
}

func startServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(conn)
	}
	
}