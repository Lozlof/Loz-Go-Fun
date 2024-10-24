package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
// This function will always return both a string and an error. 
// The error value can be nil or an actual error value if something went wrong.
func Hello(name string) (string, error) { 
// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
// If the "name" variable is an empty string, return an empty string ("") and the "empty name" error.

// Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
// Since this function always has to return both a string and an error, 
// if the function completes successfully and no errors occur, the error returned is nil.