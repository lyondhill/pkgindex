package main

import (
	"testing"
)



func TestParse(t *testing.T) {
	validTests := map[string]pkg{
		"INDEX|good|data":pkg{Command:"INDEX", Name:"good", Dependencies: []string{"data"}},
		"REMOVE|valid|": pkg{Command:"REMOVE", Name:"valid", Dependencies: []string{}},
		"QUERY|good|more,parts,exist": pkg{Command:"QUERY", Name: "good", Dependencies: []string{"more", "parts", "exist"}},
	}
	for input, request := range validTests {
		returnedRequest, err := parse(input)
		if err != nil {
			t.Errorf("error returned on valid input(%s)", input)
			continue
		}
		if returnedRequest.Command != request.Command {
			t.Errorf("command expected (%s) but received (%s)", request.Command, returnedRequest.Command)
		}

		if returnedRequest.Name != request.Name {
			t.Errorf("name expected (%s) but received (%s)", request.Name, returnedRequest.Name)
		}

		if len(returnedRequest.Dependencies) != len(request.Dependencies) {
			t.Errorf("dependency length mis match")
		}

		// no need to check the order of dependencies just the inclusion
		// ensure the expected dependencies are on the list
		for _, reqDep := range request.Dependencies {
			found := false
			for _, respDep := range returnedRequest.Dependencies {
				if reqDep == respDep {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("did not find an expected dependency in the response(%s)", reqDep)
			}
		}

		// now make sure the response doesnt have anything its not supposed to
		for _, respDep := range returnedRequest.Dependencies {
			found := false
			for _, reqDep := range request.Dependencies {
				if reqDep == respDep {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("did not find an response dependency in the expected list(%s)", respDep)
			}
		}


	}

	invalidTests := []string{
		"INDEX|to|many|parts",
		"REMOVE|notenoughparts",
		"bad|command|",
	}

	for _, invalidTest := range invalidTests {
		if _, err := parse(invalidTest); err == nil {
			t.Errorf("parse didnt error when it should have with bad input(%s)", invalidTest)
		}

	}
	
}

func TestValidCommands(t *testing.T) {
	validTests := []string{"INDEX", "REMOVE", "QUERY"}
	for _, validTest := range validTests {
		if !validCommand(validTest) {
			t.Errorf("failed with valid input(%s)", validTest)
		}
	}

	invalidTests := []string{"INDEXS", "Remove", "ADD"}
	for _, invalidTest := range invalidTests {
		if validCommand(invalidTest) {
			t.Errorf("succseded with invalid input (%s)", invalidTest)
		}
	}

}


