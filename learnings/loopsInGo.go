package learnings

import "fmt"

func LearnLoops() {
	// similar to C++
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println(i)
	}

	for {
		fmt.Println("Never ending loop!")
		break
	}
}
