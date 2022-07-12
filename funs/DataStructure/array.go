package main

import "fmt"

func main() {
	var intArr = []int{11, 12, 13}
	fmt.Println(intArr)
	fmt.Println(len(intArr))

	fmt.Println("Invoking checkPassingArrayAsValue()")
	checkPassingArrayAsValue(intArr)

	fmt.Println("Checking intArr after invoking checkPassingArrayAsValue()")
	fmt.Println(intArr)

	integerArr := intArr
	integerArr[0] = 3
	fmt.Println(integerArr)
	fmt.Println(intArr)
}

func checkPassingArrayAsValue(intgArr []int) {
	intgArr[0] = 10
	fmt.Println(intgArr)
}
