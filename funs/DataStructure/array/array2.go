package main

import "fmt"

func main() {
	days := [...]string{"Sunday", 1: "Monday", "Tuesday", 6: "Saturday", "Thursday", 5: "Friday", 3: "Wednesday", 10: "Fooday"}
	fmt.Printf("Array Length: %d\n", len(days))
	fmt.Println("Index", "  Value")
	for i, v := range days {
		fmt.Printf("  %d     %s\n", i, v)
	}

	checkArrayComparison()
}

func checkArrayComparison() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 4}
	if arr1 == arr2 {
		fmt.Println("arr1 and arr2 are equal")
	}
	if arr2 != arr3 {
		fmt.Println("arr2 and arr3 are not equal")
	}
}
