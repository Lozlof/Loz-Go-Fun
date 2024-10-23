// go env GOPATH : View environmental variable GOPATH.
// go env : View all environmental variables.

// go mod init Loz-Go-Fun : Creates a new module, initializing the go.mod file that describes it.
// go get rsc.io/quote : Changes the required version of a dependency (or adds a new dependency).
// go mod tidy : Removes unused dependencies.
// go build, go test, and other package-building commands add new dependencies to go.mod as needed.
// go run . : Compiles and the Go program in one step.

// Package declaration.
// The main package is special in Go because it's used for programs that are meant to be executable.
// Every Go program that you want to run must have a main package, and within that package, it must have a main function. 
// This is the entry point of your program, and it tells Go where to start executing code.
package main

// The import statement is used to bring in other packages from the Go standard library.
// It allows you to access functions, types, and methods that are defined in other packages. 
import "fmt" // Importing fmt package for formatted I/O.

// Imports can also be done with an import block.
// You can search for packages at https://pkg.go.dev
import (
	"math" // Importing math package for mathematical functions.
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!") // Using fmt package to print to console.
	fmt.Println(math.Sqrt(16)) // Using math package to calculate square root.
	// "quote" is the package and "Go" is the function inside the package.
    	// Meaning you need to know the source code of the package and the function, or at least know how to use them.
    	// Package info: https://pkg.go.dev/rsc.io/quote/v4#Go
    	// Function source code: https://github.com/rsc/quote/blob/v4.0.1/quote.go#L22
	fmt.Println(quote.Go())
}

// In bash, home directory:
// mkdir ~/greetings, cd greetings
// go mod init example.com/greetings, go.mod file is created in ~/greetings
// The go mod init command creates a go.mod file to track your code's dependencies. So far, the file includes only the name of your module and the Go version your code supports. But as you add dependencies, the go.mod file will list the versions your code depends on. This keeps builds reproducible and gives you direct control over which module versions to use.
// nano greetings.go, pasted the code into the file and saved.
// mkdir ~/hello, cd hello
// go mod init example.com/hello

// greetings.go 
//// package greetings
//import "fmt"
//// Hello returns a greeting for the named person.
//func Hello(name string) string {
    //// Return a greeting that embeds the name in a message.
    //message := fmt.Sprintf("Hi, %v. Welcome!", name)
    //return message
//}
// This function takes a name parameter whose type is string. The function also returns a string.
// In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.
// Function name: "Hello"
// Parameter type: "string"
// Return type: "string"
// In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type). 
// Taking the long way, you might have written this as:
// var message string
// message = fmt.Sprintf("Hi, %v. Welcome!", name)
