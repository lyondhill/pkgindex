package main

import (
	"bufio"
	"fmt"
	"io"
)

type handleFunc func(*pkg) error

var handlers = map[string]handleFunc{}

// setup the handlers so it is more expandable
func init() {
	handlers["INDEX"] = index
	handlers["QUERY"] = query
	handlers["REMOVE"] = remove
}

// handleConn handles connections as they come in
func handleConn(conn io.ReadWriteCloser) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		fmt.Fprintln(conn, handleRequest(scanner.Text()))
	}

}

// takes a request string and responds with the result from the pkg system
func handleRequest(request string) string {
	// parse the request
	pkg, err := parse(request)
	if err != nil {
		return "ERROR"
	}

	// pull the function out of the handlers
	// this should never fail because the parser will error before this
	fn := handlers[pkg.Command]

	// execute the function and check for errors
	err = fn(pkg)
	if err != nil {
		return "FAIL"
	}

	return "OK"
}
