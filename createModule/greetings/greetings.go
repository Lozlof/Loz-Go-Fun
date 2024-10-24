package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time" // Had to import time in order for rand.Seed to work.
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

	// Creates a message using a random format.
	// Calls the Sprintf function from the fmt package.
	// Sprintf is used to format a string without printing it.
	// The first argument to fmt.Sprintf is a format string, and the subsequent arguments provide the values to insert into the format string.
	// Sprintf formats according to a format specifier and returns the resulting string.
	// %v is a standardized placeholder in Go, and it is part of Go's powerful formatting system provided by the fmt package.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}
// Since this function always has to return both a string and an error, 
// if the function completes successfully and no errors occur, the error returned is nil.

func Hellos(names []string) (map[string]string, error) {
	// Leave off at: https://go.dev/doc/tutorial/greetings-multiple-people
}

// Function declaration.
// Name: randomFormat
// Note: That randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported). 
// Parameters: The empty parentheses () indicate that this function does not take any parameters. It doesn’t expect any inputs when called.
// Return Type: Will return a value type of string.
func randomFormat() string {
	// Declaring and initializing the formats variable.
	// The type of formats is []string meaning that it is a slice of strings.
	// formats contains three string values.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Greetings! %v!",
		"Good Day! %v, how are you?",
	}

	rand.Seed(time.Now().UnixNano()) // Have to seed for rand.Intn to work.

	return formats[rand.Intn(len(formats))]
	// From: https://pkg.go.dev/math/rand#Intn
	// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n) from the default Source. It panics if n <= 0.
	// func Intn(n int) int
	// The Intn function takes n an an argument.
	// Returns an integer value.
	// rand.Intn(n) generates a random integer in the half-open interval [0, n). This means:
	// The generated number is always greater than or equal to 0.
	// The generated number is always less than n (i.e., n is exclusive).
	// So if n is 10, it will return a random number between 0 and 9 (inclusive).
	// Instead of passing a static integer to rand.Intn, you are passing the length of formats.
}	
