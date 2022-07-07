// To check the init function

package main

import "fmt" // 1. importing pkgs are resolved

var a, b int = valueOfa(), 5 // 2. pkg-level variables and dependecies are resolved

// 3. init functions are invoked
func init() {
	fmt.Printf("a+b: %d\n", a+b)
}

func init() {
	c := 30
	a = c
	b = 50

	fmt.Println(a)
	fmt.Println(b)
}

func valueOfa() int {
	return 2
}

// 4. finally main function is invoked
func main() {
	fmt.Println("Init function check")
}
