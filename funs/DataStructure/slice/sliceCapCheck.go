package main

import "fmt"

func main() {
	fmt.Println("#slice capacity check")
	sliceCapacityCheck()
}

func sliceCapacityCheck() {
	xSlice := []int{6, 7}
	fmt.Printf("len(xSlice): %d\n", len(xSlice))
	fmt.Printf("cap(xSlice): %d\n", cap(xSlice))
	ySlice := append(xSlice, 9, 10, 11)
	fmt.Printf("len(ySlice): %d\n", len(ySlice))
	fmt.Printf("cap(ySlice): %d\n", cap(ySlice)) // len(ySlice) = 5 and cap(ySlice) = 6. Why cap(ySlice) is not 5?
	fmt.Println()

	tSlice := []int{1, 2}
	fmt.Printf("len(tSlice): %d\n", len(tSlice))
	fmt.Printf("cap(tSlice): %d\n", cap(tSlice))
	zSlice := append(tSlice, 9, 10, 11, 12)
	fmt.Printf("len(zSlice): %d\n", len(zSlice))
	fmt.Printf("cap(zSlice): %d\n", cap(zSlice))

	fmt.Println()
	aSlice := []int{3, 4, 9}
	fmt.Printf("len(aSlice): %d\n", len(aSlice))
	fmt.Printf("cap(aSlice): %d\n", cap(aSlice))
	fSlice := []int{0, 6, 7}
	bSlice := append(aSlice, fSlice...)
	fmt.Printf("len(bSlice): %d\n", len(bSlice))
	fmt.Printf("cap(bSlice): %d\n", cap(bSlice))
}
