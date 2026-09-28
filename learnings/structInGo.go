package learnings

import "fmt"

// Fields whose names start with a capital letter are exported and can be accessed from other packages.
type Person struct {
	Name string
	Age  int
}

// Go has no special constructor syntax. They're just ordinary funnctions.

/*
Returning a pointer is useful when:

1. you want shared identity,
2. expect the value to be modified through the pointer,
4. or want to avoid copying a large struct.
*/
func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

// Structs are value types, so assigning one copies its fields

func LearnStruct() {
	person := Person{Name: "Satyajit", Age: 28}
	fmt.Println(person)

	person.Age = 29
	fmt.Println(person)

	p := &person
	p.Age = 28 // Go automatically dereferences p here
	fmt.Println(person)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}

// This is a method defined on struct Person
// Method name is Greet. It takes no args.
// Method doesn't return any value.
/*
	A receiver is similar to a function parameter,
	but it makes the function a method of that type.
*/
func (p Person) Greet() {
	fmt.Println("Hello,", p.Name)
}
