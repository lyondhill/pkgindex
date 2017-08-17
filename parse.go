package main

import (
	"errors"
	"strings"
)

const INDEX = "INDEX"

// List of all available commands the server can handle
var validCommands = [3]string{"INDEX", "REMOVE", "QUERY"}

// error definition of all error request
var invalidRequest = errors.New("invalid request")

func parse(input string) (*pkg, error) {
	// clean any newlines that may be present
	input = strings.Replace(input, "\n", "", -1)

	// split the request into its subjective parts
	requestParts := strings.Split(input, "|")

	// make sure we atleast have the correct number of parts
	if len(requestParts) != 3 {
		return nil, invalidRequest
	}

	// retrieve the dependencies from the request if there are any
	dependencies := []string{}
	if len(requestParts[2]) > 0 {
		dependencies = strings.Split(requestParts[2], ",")
	}

	// return the resulting request object
	response := pkg{
		Command:       requestParts[0],
		Name:          requestParts[1],
		Dependencies:  dependencies,
		DependantPkgs: []*pkg{},
	}

	// check to make sure the command is a valid command
	if !validCommand(response.Command) {
		return nil, invalidRequest
	}

	return &response, nil
}

// Ensure that the command given is valid
func validCommand(command string) bool {
	for _, validCommand := range validCommands {
		if command == validCommand {
			return true
		}
	}
	return false
}
