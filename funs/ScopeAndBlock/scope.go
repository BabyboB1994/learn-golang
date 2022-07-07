package main

import "fmt"

var w string = "hello"

func main() {
	fmt.Println(w)
	r := 9

	w, r := check()
	fmt.Println(w)
	fmt.Println(r)
}

func check() (string, int) {
	return "bingo", 3
}
