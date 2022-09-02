package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6, 7, 8}

	fmt.Println("Before copy")
	fmt.Println("slice1--> ", slice1)
	fmt.Println("slice2--> ", slice2)
	fmt.Println("After copy")
	fmt.Println("no of elements copied: ", copy(slice1, slice2[3:]))
	fmt.Println("slice1--> ", slice1)
	fmt.Println("slice2--> ", slice2)

	// sliceCheck := slice1[3:] // this returns a non-nil zero length slice
	// slcheader := (*reflect.SliceHeader)(unsafe.Pointer(&sliceCheck))
	// fmt.Printf("s7Header: %+v\n", slcheader)
}
