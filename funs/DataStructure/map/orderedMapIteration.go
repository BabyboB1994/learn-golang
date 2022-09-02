package main

import (
	"fmt"
	"sort"
)

var mapCheck = map[int]string{
	1: "apple",
	2: "ball",
	3: "icecream",
	4: "den",
	5: "elephant",
}

func main() {
	for i, v := range mapCheck {
		fmt.Println(i, " ", v)
	}
	accessMapInOrder(mapCheck)
}

func accessMapInOrder(mapCheck map[int]string) {
	var keys []int
	for i := range mapCheck {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	for _, key := range keys {
		fmt.Println(key, " ", mapCheck[key])
	}
}

// When iterating over a map with a range loop, the iteration order is not specified
// and is not guaranteed to be the same from one iteration to the next
/*
go run map2.go
1   apple
2   ball
3   icecream
4   den
5   elephant

go run map2.go
4   den
5   elephant
1   apple
2   ball
3   icecream
*/
