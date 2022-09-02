package main

import (
	"fmt"
	"reflect"
)

type ss struct{}

func main() {
	var st ss
	fmt.Println(st)
	fmt.Println(reflect.TypeOf(st))
}
