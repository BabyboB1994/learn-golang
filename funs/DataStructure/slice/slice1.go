package main

import "fmt"

func main() {
	var buffer = []int{1, 2, 3, 4, 5}
	fmt.Println("buffer before invoking the func: ", buffer)
	fmt.Println("Address of []buffer elements:")
	fmt.Printf("&buffer[0]: %p\n", &buffer[0])
	fmt.Printf("&buffer[1]: %p\n", &buffer[1])
	fmt.Printf("&buffer[2]: %p\n", &buffer[2])
	fmt.Printf("&buffer[3]: %p\n", &buffer[3])
	fmt.Printf("&buffer[4]: %p\n\n", &buffer[4])

	slice := buffer[:4]
	fmt.Println("slice before invoking the func: ", slice)
	fmt.Println("Length: ", len(slice))
	fmt.Println("Capacity: ", cap(slice))
	fmt.Println("Address of []slice elements:")
	fmt.Printf("&slice[0]: %p\n", &slice[0])
	fmt.Printf("&slice[1]: %p\n", &slice[1])
	fmt.Printf("&slice[2]: %p\n", &slice[2])
	fmt.Printf("&slice[3]: %p\n\n", &slice[3])

	incrementEachElement(slice) // copy of slice is passed here

	// although a copy of slice is passed, it contains a ptr to underlying array. Hence change(s) made will reflect
	fmt.Println("After invoking the func")
	fmt.Println("[]slice: ", slice)
	fmt.Println("[]buffer: ", buffer)
}

func incrementEachElement(slice []int) {
	for i := range slice {
		slice[i] += 1
	}
}

// output
/*
buffer before invoking the func:  [1 2 3 4 5]
Address of []buffer elements:
&buffer[0]: 0xc00001c0f0
&buffer[1]: 0xc00001c0f8
&buffer[2]: 0xc00001c100
&buffer[3]: 0xc00001c108
&buffer[4]: 0xc00001c110

slice before invoking the func:  [1 2 3 4]
Length:  4
Capacity:  5
Address of []slice elements:
&slice[0]: 0xc00001c0f0
&slice[1]: 0xc00001c0f8
&slice[2]: 0xc00001c100
&slice[3]: 0xc00001c108

After invoking the func
[]slice:  [2 3 4 5]
[]buffer:  [2 3 4 5 5]
*/
