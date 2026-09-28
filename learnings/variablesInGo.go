package learnings

import "fmt"

func LearnVariables() {
	// not possible: var a int, b float64
	var c, d string
	var a int
	var b float64
	a = 1
	b = 2.0
	c = "Hello"
	d = "World"

	// Printing the values of the variables
	fmt.Printf("Value of a: %d, Type of variable: %T\n", a, a)
	fmt.Printf("Value of b: %f, Type of variable: %T\n", b, b)
	fmt.Printf("Value of c: %s, Type of variable: %T\n", c, c)
	fmt.Printf("Value of d: %s, Type of variable: %T\n", d, d)

	x := 10
	// x := 10 equivalent to 'var x int = 10'
	fmt.Printf("Value of x: %d, Type of variable: %T\n", x, x)
	y := x + a
	fmt.Printf("Value of y: %d, Type of variable: %T\n", y, y)
}
