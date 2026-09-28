package learnings

import (
	"fmt"
	"unicode/utf8"
)

func LearnStrings() {
	for i, c := range "go" {
		fmt.Println(i, c)
		// A Go string is a read-only slice of bytes.
		// strings are equivalent to []byte
		// In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.
	}

	var a string
	a = "Satyajit"
	// a[1] = 'c' fails because Go strings are immutable.

	text := []byte(a) // mutatable byte slice
	text[1] = 'c'
	a = string(text)

	fmt.Println(a)
	fmt.Println("Rune count:", utf8.RuneCountInString(a))

	for idx, runeValue := range a {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// anonymous function
	examineRune := func(r rune) {

		switch r {
		case 't':
			fmt.Println("found tee")
		case 'ส':
			fmt.Println("found so sua")
		}
	}

	const s = "สวัสดี"
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		// utf8.DecodeRuneInString(s[i:]) decodes the next Unicode code point and returns its byte width.
		// i is a byte offset, so the loop advances by w, the width of the rune just decoded.
		// For this Thai string, each rune uses 3 bytes.

		fmt.Printf("%#U starts at %d\n", runeValue, i)
		// %#U prints the rune as a Unicode code point, and its starting byte offset.
		w = width
		examineRune(runeValue)
	}
}
