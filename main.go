package main

import "fmt"

func main() {
	// You can declare the variable like so:
	// var input string
	// input = "The quick brown fox jumped over the lazy dog"

	// But most of the time you see the shorthand:
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}

func Reverse(s string) string {
	// You can declare a variable this way:
	// var b []byte
	// b = []byte(s)

	// OR  you can use the shortcut like so:
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
