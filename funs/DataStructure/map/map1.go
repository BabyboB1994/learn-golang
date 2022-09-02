package main

import (
	"fmt"
)

func main() {
	// map syntax : map[KeyType]ValueType

	var months map[string]int // creates nil map
	fmt.Println(months)
	if months == nil {
		fmt.Println("months is nil")
	}

	points := map[string]float64{}
	fmt.Println(points)
	if points != nil {
		fmt.Println("points is not nil")
	}

	years := make(map[int]int)
	fmt.Println(years)
	if years != nil {
		fmt.Println("years is not nil")
	}

	students := map[string]int{
		"Bob":    1,
		"Keith":  2,
		"Mogran": 3,
	}
	students["Darryl"] = 7
	fmt.Println(students)

	days := make(map[int]string, 3)
	days[0] = "Sunday"
	days[1] = "Monday"
	days[2] = "Tuesday"
	days[3] = "Wednesday"
	days[4] = "Thursday"
	days[5] = "Friday"
	days[6] = "Saturday"
	fmt.Println(days)
	fmt.Println(days[6])
	fmt.Println("length of days: ", len(days))

	checkMap := make(map[string]int)
	checkMap["abc"] = 99
	fmt.Println(checkMap["abc"])
	x, d := checkMap["xyz"]
	fmt.Println(x, d)
	y, k := checkMap["asd"]
	fmt.Println(y, k)
	r, t := checkMap["abc"]
	fmt.Println(r, t)
}
