package learnings

import (
	"fmt"
	"maps"
)

func LearnMaps() {
	m := make(map[string]int)
	m["key-1"] = 1
	m["key-2"] = 2

	fmt.Println("map:", m)
	delete(m, "key-2")
	fmt.Println("map:", m)
	clear(m)
	fmt.Println("map:", m)

	_, res := m["key-2"]
	fmt.Println("Key Exists:", res)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	//Equal can be used for arrays and slices
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
