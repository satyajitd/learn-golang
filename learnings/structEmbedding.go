package learnings

import "fmt"

type KnowHuman interface {
	printAddress()
}

type Address struct {
	City    string
	Pincode int
}

// Method for struct Address
func (a Address) printAddress() {
	fmt.Println(a.City, a.Pincode)
}

/*
Go does not have class-based inheritance.
Instead, it uses composition: a struct can embed another type, which promotes that type’s fields and methods.
*/
type Human struct {
	Name    string
	Address // embedded field; its field name is Address
}

func LearnStructEmbedding() {
	human := Human{
		Name:    "Satyajit",
		Address: Address{City: "Bengaluru", Pincode: 560075},
	}

	// Embedding promotes the inner struct’s fields, so human.City is shorthand for human.Address.City.
	fmt.Println(human.City)
	// If multiple embedded types expose the same field or method name, use the explicit path, such as human.Address.City.
	fmt.Println(human.Address.City)

	// Methods are promoted too: if Address has a method, you can call it through Human.
	// But embedding is composition, not inheritance
	human.printAddress()

	// Human satisfies the interface because Human gets the method printAddress through embedding
	// Embedded Address promotes Address.printAddress onto Human
	var knowHuman KnowHuman = human
	knowHuman.printAddress()
}
