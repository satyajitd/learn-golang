package learnings

import "fmt"

func LearnConditions() {
	var n int = 5
	if n%2 == 0 {
		fmt.Println("'n' is an even number")
	} else {
		fmt.Println("'n' is an odd number")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
