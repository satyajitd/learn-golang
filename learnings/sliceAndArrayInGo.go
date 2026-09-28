package learnings

import "fmt"

func LearnSliceAndArray() {
	// Array
	var a [5]int
	fmt.Println(a)
	b := [...]int{1, 2, 3, 4}
	fmt.Println(len(b), b)

	twoD := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println(twoD)

	// Slice
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	s = append(s, "e", "f")
	fmt.Println(s[:1])

	c := make([]string, len(s))
	copy(c, s)

	twoDSlice := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoDSlice[i] = make([]int, innerLen)
		for j := range innerLen {
			twoDSlice[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoDSlice)
}
