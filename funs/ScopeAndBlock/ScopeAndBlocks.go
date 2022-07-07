package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Hostname())
	justCheck()
}

func sampleFunc() {
	fmt.Println("Hello from sampleFunc()")
}
