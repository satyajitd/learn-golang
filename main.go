// Code files are generated with extension .go
// To run Go code, execute 'go run <file_name>.go' in the terminal
// To build Go code, execute 'go build <file_name>.go' in the terminal

package main

// What is the purpose of the 'main' package in Go?
// The 'main' package in Go serves a special purpose: it defines a standalone executable program.
// When you compile and run a Go program, the Go runtime looks for the 'main' package and specifically the 'main' function within that package to start execution.
// The 'main' function is the entry point of the program, and it is where the program begins its execution.
// Without a 'main' package and a 'main' function, the Go code cannot be executed as a standalone application.

import (
	learnings "learn-go/learnings"
)

func main() {
	learnings.LearnVariables()
	learnings.LearnOperations()
	learnings.LearnConditions()
	learnings.LearnLoops()
	learnings.LearnSwitch()
	learnings.LearnSliceAndArray()
	learnings.LearnMaps()
	learnings.LearnFunctions()
	learnings.LearnStrings()
	learnings.LearnPointers()
	learnings.LearnStruct()
	learnings.LearnInterface()
	learnings.LearnEnums()
	learnings.LearnStructEmbedding()
}
