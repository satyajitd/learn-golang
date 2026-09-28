package learnings

import "fmt"

// What fmt package is used for in Go?
// The 'fmt' package in Go is a standard library package that provides formatted I/O (input/output) functions.
// It is commonly used for printing output to the console and reading input from the user.
// The 'fmt' package includes functions like Println, Printf, Sprintf, Scan, and many others that allow developers to format strings, print values, and read input in a convenient way.
// In this code snippet, the 'fmt' package is used to print "Hello World" to the console using the Println function.

func LearnOperations() {
	// String concatenation in Go using the '+' operator
	fmt.Println("My second" + " lesson")
	// Performing arithmetic operations in Go
	fmt.Println("1 + 2 = ", 1+2)
	fmt.Printf("Type of 1 + 2: %T\n", 1+2)
	fmt.Println("7.0 / 2.0 = ", 7.0/2.0)
	fmt.Printf("Type of 7.0 / 2.0: %T\n", 7.0/2.0)
	fmt.Println("7.0 * 2 = ", 7.0*2)
	fmt.Printf("Type of 7.0 * 2: %T\n", 7.0*2)
	fmt.Println("7 * 2.0 = ", 7*2.0)
	fmt.Printf("Type of 7 * 2.0: %T\n", 7*2.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
}
