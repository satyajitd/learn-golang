package learnings

import "fmt"

func plus(a int, b int) (int, int, int) {
	return a + b, a, b
}

// function overloading not allowed in Golang
func sum(a, b int) int {
	return a + b
}

// variadic function
func variadicSum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

/*
A closure in Go is a function value that can use variables from the scope where it was created.
It keeps access to those variables even after the outer function returns.
*/
func makeCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func LearnFunctions() {
	// slice
	values := []int{1, 2, 3}
	variadicSum(values...)
	variadicSum(1, 3, 4, 5)

	counter := makeCounter()
	fmt.Println(counter())

	// anonymous function
	helloWorld := func() {
		fmt.Println("Hello, World!")
	}
	helloWorld()
}
