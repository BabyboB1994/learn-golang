package main

import "fmt"

type month int
type integer = int // here integer is alias of int

func main() {
	var assgnCheck integer
	var er month = 2

	var currentMonth int = er // this cannot be done even though the uderlying type of month is a integer

	i := 9
	assgnCheck = i
	fmt.Println(assgnCheck)
}
