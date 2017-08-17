package main

import (
	"fmt"
	"net"
	"testing"
)

func TestStartServer(t *testing.T) {
	go startServer(1234)

	conn, _ := net.Dial("tcp", "localhost:1234")
	fmt.Fprintln(conn, "QUERY|hi|")
	var result string
	fmt.Fscanln(conn, &result)
	if result != "fail" {
		t.Errorf("recieved invalid result: %s", result)
	}
	conn.Close()
}
