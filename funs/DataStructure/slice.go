package main

import "fmt"

func main() {
	var buffer []byte = []byte{1, 2, 3, 4, 5}
	slice := buffer[:4]

	fmt.Println("before", slice)
	AddOneToEachElement(slice) // copy of slice is passed here
	fmt.Println("after", slice)
}

func AddOneToEachElement(slice []byte) {
	slice = slice[1:3]
}
