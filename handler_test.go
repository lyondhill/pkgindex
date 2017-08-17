package main

import (
	"fmt"
	"net"
	"testing"
)

func TestHandlerConn(t *testing.T) {
	testCases := [][]string{
		[]string{"hi|this|shouldnt|work", "ERROR"},
		[]string{"hi|thisisbadtoo|", "ERROR"},
		[]string{"INDEX|my|", "OK"},
		[]string{"QUERY|my|", "OK"},
		[]string{"QUERY|this should fail", "ERROR"},
		[]string{"INDEX|a|bad,dep", "FAIL"},
		[]string{"REMOVE|my|", "OK"},
	}

	client, server := net.Pipe()
	go handleConn(server)

	for i := 0; i < len(testCases); i++ {
		fmt.Fprintln(client, testCases[i][0])

		var result string
		fmt.Fscanln(client, &result)

		if result != testCases[i][1] {
			t.Errorf("invalid response for (%#v): expected(%#v) != recieved(%#v)", testCases[i][0], testCases[i][1], result)
		}
	}
	client.Close()

}
