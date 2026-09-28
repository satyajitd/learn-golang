package learnings

import "fmt"

func LearnPointers() {
	// pointer
	p := new(42)
	fmt.Println(*p)

	var ptr *int
	ptr = new(42)
	fmt.Println(*ptr)
}
