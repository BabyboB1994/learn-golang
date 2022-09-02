package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	var arr2 = arr1
	fmt.Println("Before")
	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)
	arr2[0] = 12
	fmt.Println("After")
	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)

	var slice1 = []int{2, 3, 4}
	var slice2 = slice1
	fmt.Println("Before")
	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)
	slice2[0] = 13
	fmt.Println("After")
	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)
}

//output
/*
Before
arr1:  [1 2 3]
arr2:  [1 2 3]
After
arr1:  [1 2 3]
arr2:  [12 2 3]
*/
