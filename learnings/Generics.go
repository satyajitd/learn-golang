package learnings

/*
Go supports generics: functions and types can take type parameters, so you can write reusable code while preserving compile-time type checking.
*/

import "fmt"

// A generic function declares type parameters in square brackets
// T is a type parameter, and any is a constraint meaning it can be any type.
func First[T any](values []T) T {
	return values[0]
}

func LearnGenerics() {
	names := []string{"Ada", "Lin"} // slice
	fmt.Println(First(names))       // T is inferred as string
}

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	_ = SlicesIndex(s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
