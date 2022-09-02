package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)
	fmt.Println("Length of arr: ", len(arr))
	fmt.Println("Capacity of arr: ", cap(arr))
	fmt.Println()

	slice := arr[3:7]
	fmt.Println(slice)
	fmt.Println("Length of slice: ", len(slice))
	fmt.Println("Capacity of slice: ", cap(slice))
	fmt.Println()

	slice2 := slice[:7]
	fmt.Println("slice2: ", slice2)
	fmt.Printf("len(slice2) : %d\n", len(slice2))
	fmt.Printf("cap(slice2) : %d\n", cap(slice2))
	reverse(slice2)
	fmt.Println()
	fmt.Println(slice2)

	fmt.Println()
}

func reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
