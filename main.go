// Package declaration.
// The main package is special in Go because it's used for programs that are meant to be executable.
// Every Go program that you want to run must have a main package, and within that package, it must have a main function. 
// This is the entry point of your program, and it tells Go where to start executing code.
package main

// The import statement is used to bring in other packages from the Go standard library.
// It allows you to access functions, types, and methods that are defined in other packages. 
import "fmt" // Importing fmt package for formatted I/O.

// Imports can also be done with an import block.
import (
	"math" // Importing math package for mathematical functions.
)

func main() {
	fmt.Println("Hello, World!") // Using fmt package to print to console.
	fmt.Println(math.Sqrt(16)) // Using math package to calculate square root.
}
