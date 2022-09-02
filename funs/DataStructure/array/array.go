package main

import "fmt"

func main() {
	var emptArr = [...]int{}
	fmt.Println(emptArr)
	fmt.Println(cap(emptArr))
	fmt.Println()

	var intArr = [...]int{11, 12, 13}
	fmt.Printf("%p\n", &intArr)
	fmt.Printf("Elements in intArr: %v\n", intArr)
	fmt.Printf("Type of intArr: %T\n", intArr)
	fmt.Printf("Length of intArr: %d\n", len(intArr))
	//fmt.Println("Lenght of intArr: "+string(len(intArr)))
	fmt.Printf("Addr. of intArr[0] : %p\n", &intArr[0])
	fmt.Printf("Addr. of intArr[1] : %p\n", &intArr[1])
	fmt.Printf("Addr. of intArr[2] : %p\n", &intArr[2])
	fmt.Println()

	fmt.Println("Invoking checkPassingArrayAsValue()")
	checkPassingArrayAsValue(intArr)
	fmt.Println()

	fmt.Println("Checking intArr after invoking checkPassingArrayAsValue()")
	fmt.Println(intArr)
	fmt.Println()

	integerArr := intArr
	integerArr[0] = 3
	fmt.Printf("Elements in integerArr: %v\n", integerArr)
	fmt.Printf("Addr. integerArr[0] : %x\n", &integerArr[0])
	fmt.Printf("Addr. integerArr[1] : %x\n", &integerArr[1])
	fmt.Printf("Addr. integerArr[2] : %x\n", &integerArr[2])

	fmt.Printf("Elemets in intArr: %v\n", intArr)
	fmt.Printf("Addr. intArr[0] : %x\n", &intArr[0])
	fmt.Printf("Addr. intArr[1] : %x\n", &intArr[1])
	fmt.Printf("Addr. intArr[2] : %x\n", &intArr[2])

}

func checkPassingArrayAsValue(intgArr [3]int) {
	fmt.Printf("Addr. intgArr[0] : %x\n", &intgArr[0])
	fmt.Printf("Addr. intgArr[1] : %x\n", &intgArr[1])
	fmt.Printf("Addr. intgArr[2] : %x\n", &intgArr[2])
	intgArr[0] = 99
	intgArr[1] = 98
	intgArr[2] = 97
	fmt.Printf("Elements in intgArr: %v\n", intgArr)
}

// sample output
/*
[]
0

0xc0000140c0
Elements in intArr: [11 12 13]
Type of intArr: [3]int
Length of intArr: 3
Addr. of intArr[0] : 0xc0000140c0
Addr. of intArr[1] : 0xc0000140c8
Addr. of intArr[2] : 0xc0000140d0

Invoking checkPassingArrayAsValue()
Addr. intgArr[0] : c000014108
Addr. intgArr[1] : c000014110
Addr. intgArr[2] : c000014118
Elements in intgArr: [99 98 97]

Checking intArr after invoking checkPassingArrayAsValue()
[11 12 13]

Elements in integerArr: [3 12 13]
Addr. integerArr[0] : c000014150
Addr. integerArr[1] : c000014158
Addr. integerArr[2] : c000014160
Elemets in intArr: [11 12 13]
Addr. intArr[0] : c0000140c0
Addr. intArr[1] : c0000140c8
Addr. intArr[2] : c0000140d0
*/
