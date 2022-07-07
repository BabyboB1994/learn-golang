package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	println("Hello from GO!!")

	var xyz = 0x_67_7a_2f_cc_40_c6
	fmt.Println(reflect.TypeOf(xyz))
	fmt.Println(xyz)

	// just to check the max size of a int
	fmt.Println(math.MaxInt)
	fmt.Println(math.MaxInt64)

	// fmt.Println(math.MaxUint)

	const uinteger uint = math.MaxUint
	const uintgr = math.MaxUint /// Why math.MaxUint is an untyped int, not untyped uint?
	fmt.Println(uinteger)

	// from Rob Pike's article about GO constants
	const MaxUint = ^uint(0)
	fmt.Printf("%v\n", MaxUint)
}
