package main

import (
	"fmt"
)

func main() {
	//s1 := "tom and jerry"
	s2 := "hello"
	var s3 = s2[0]

	//s2[0] = 'K'

	fmt.Println(s2, "|", s3)

	fmt.Printf("Addres of s2: %p\n", &s2)
	fmt.Printf("Addres of s3: %p\n", &s3)

	// s2 += " world khkhrghfhg hhroietoirhoighreiotireglregiehoieri"
	// fmt.Println(s2, "|", s3)

	// fmt.Printf("Addres of s2: %p\n", &s2)
	// fmt.Printf("Addres of s3: %p\n", &s3)
}
