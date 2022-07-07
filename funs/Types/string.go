package main

import "fmt"

func main() {
	s := "foo-labs"
	fmt.Println(len(s))
	fmt.Println(s[3:7])
	fmt.Println(s[0:])
	fmt.Println(s[:8])

	/*The i-th byte of a string is not necessarily the i-th character of a string, because the UTF-8 encoding of a non-ASCII code point requires two or more bytes.*/
	s2 := "ഇ" // "ഇ" is a malayalam character which requires 3 bytes
	fmt.Println(len(s2))

	fmt.Println(s2[0:1]) // this prints a junk char

	// String concat
	fmt.Println(s + s2)
}
