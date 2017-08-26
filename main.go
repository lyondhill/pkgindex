// Package Index builds a indexing list with dependencies.
//
// Design
// Package Index is designed to be simple and easy to understand. It uses a in-memory data storage and runs on a single node with no external dependencies. The data is stored in a `map[string]*pkg` where pkg is:
//
// type pkg struct {
// 	Command       string
// 	Name          string
// 	Dependencies  []string
// 	DependantPkgs []*pkg
// }
//
// By creating a data structure in this way it makes the primary functions very quick.
// When inserting data we need to ensure the dependencies of this package are all available, this is done quickly
// by indexing the map. When removing a package we need to guarantee the package we are removing does not have
// anything depending upon it, and this is done quickly by having the Dependant packages referenced.
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on (default 8080)")
	flag.Parse() // initiate the server

	startServer(*port)
}

func startServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go handleConn(conn)
	}

}
