package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	slice := []int{1, 2}
	fmt.Println(slice)
	fmt.Println("length: ", len(slice))
	fmt.Println("Capacity: ", cap(slice))
	fmt.Printf("&slice: %p\n", &slice)
	fmt.Printf("&slice[0]: %p\n", &slice[0])
	fmt.Printf("&slice[1]: %p\n\n", &slice[1])
	slcheader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Printf("s7Header: %+v\n", slcheader)

	slice2 := append(slice, 3, 4, 5) // new underlying array will be allocated as slice doesn't have sufficient memory
	fmt.Println(slice2)
	fmt.Println("length: ", len(slice2))
	fmt.Println("Capacity: ", cap(slice2))
	fmt.Printf("&slice2: %p\n", &slice2)
	fmt.Printf("&slice2[0]: %p\n", &slice2[0])
	fmt.Printf("&slice2[1]: %p\n\n", &slice2[1])
	slcheader2 := (*reflect.SliceHeader)(unsafe.Pointer(&slice2))
	fmt.Printf("s7Header: %+v\n", slcheader2)

	//slice = append(slice, 6)
	fmt.Println(slice)
	fmt.Println("length: ", len(slice))
	fmt.Println("Capacity: ", cap(slice))
	fmt.Printf("&slice: %p\n", &slice)
	fmt.Printf("&slice[0]: %p\n", &slice[0])
	fmt.Printf("&slice[1]: %p\n", &slice[1])
}
