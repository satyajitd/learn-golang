module learn-go

go 1.26

// go.mod defines your Go module: the project root and the settings Go uses to build it.
// It is not the same as a package; one module can contain several packages in different directories.

// module learn-go sets the module’s import-path prefix. 
// For example, the package in the learnings directory has the import path learn-go/learnings.

// go 1.26 tells Go which language version the module targets. 
// That matters because your code uses features such as new(42) that require Go 1.26.

// Go also records external dependencies in go.mod when you add them. 
// `go run .` reads this file to build the module from its root.