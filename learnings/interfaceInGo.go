package learnings

import (
	"fmt"
	"math"
)

/*
	Q: What is interface?

	A Go interface describes behavior: it lists methods a type must have.
	A type satisfies an interface automatically by implementing those methods;
	there’s no implements declaration.
*/

/*
	Q: Is it necessary to use struct to implement interface?

	No, any defined type can satisfy an interface by having its required methods
	any defined type can satisfy an interface by having its required methods
*/

/*
	Q: What is type?

	A type in Go tells the compiler what kind of value a variable can hold and what operations are allowed on it.
	Go is statically typed, so a variable’s type is known at compile time.

	Types can describe grouped data (struct), required behavior (interface), or collections such as slices and maps.
	A type helps Go catch mismatches, such as trying to assign a string to an int.
*/

/*
	Q: What is the difference between type and var?

	var declares a variable: a place to store a value
	type declares a type: a definition that describes what values are allowed and how they behave

	type Age int
	var myAge Age = 28

	myAge is a variable of type Age
*/

type shape interface {
	area() float64
	perimeter() float64
}

type rectange struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

// A type implements an interface by defining all the methods the interface requires, with matching names and signatures.
// Go recognizes this automatically; there’s no implements keyword.

func (r rectange) area() float64 {
	return r.height * r.width
}

func (r rectange) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

// Generic function which accepts any value whose type satisfies the shape interface
func measureOfShape(s shape) {
	fmt.Println(s)
	fmt.Printf("Area: %f\n", s.area())
	fmt.Printf("Perimeter: %f\n", s.perimeter())
}

func isCircle(s shape) {
	// s.(Circle) expression is a type assertion
	// ok tells whether the assertion succeeded or not
	// c holds values
	if c, ok := s.(circle); ok == true {
		fmt.Printf("Circle radius: %f", c.radius)
	}
}

func LearnInterface() {
	rectangle := rectange{width: 10, height: 5}
	circle := circle{radius: 3}

	fmt.Println("\nRectangle:")
	measureOfShape(rectangle)
	isCircle(rectangle)

	fmt.Println("\nCircle:")
	measureOfShape(circle)
	isCircle(circle)
	fmt.Println()
}
