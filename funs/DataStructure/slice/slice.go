package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var s2 []int // nil slice. This can also be written as s2 = []int(nil)
	fmt.Println("s2 : ", s2)
	fmt.Println("len(s2) : ", len(s2))
	fmt.Println("cap(s2) : ", cap(s2))
	fmt.Println("s2 == nil? : ", s2 == nil)
	s2Header := (*reflect.SliceHeader)(unsafe.Pointer(&s2))
	fmt.Printf("%+v\n\n", s2Header)

	s1 := []int{} //non-nil zero length slice
	fmt.Println("s1 : ", s1)
	fmt.Println("len(s1) : ", len(s1))
	fmt.Println("cap(s1) : ", cap(s1))
	fmt.Println("s1 == nil? : ", s1 == nil) //false
	s1Header := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	fmt.Printf("%+v\n\n", s1Header)

	fmt.Println("creating slice using built-in 'make' function")

	s3 := make([]int, 0)
	fmt.Println("s3: ", s3)
	fmt.Println("len(s3): ", len(s3))
	fmt.Println("cap(s3): ", cap(s3))
	fmt.Println("s3 == nil? :", s3 == nil) //non-nil zero length slice
	s3Header := (*reflect.SliceHeader)(unsafe.Pointer(&s3))
	fmt.Printf("%+v\n\n", s3Header)

	s4 := make([]int, 2, 5)
	fmt.Println("s4: ", s4)
	fmt.Println("len(s4): ", len(s4))
	fmt.Println("cap(s4): ", cap(s4))
	s4Header := (*reflect.SliceHeader)(unsafe.Pointer(&s4))
	fmt.Printf("%+v\n\n", s4Header)

	s5 := make([]int, 10)[:5] //creates an unnamed array of length 10 and returns a slice of length 5
	fmt.Println("s5: ", s5)
	fmt.Println("len(s5): ", len(s5))
	fmt.Println("cap(s5): ", cap(s5))
	s5Header := (*reflect.SliceHeader)(unsafe.Pointer(&s5))
	fmt.Printf("%+v\n", s5Header)
	fmt.Println()

	s6 := []int{1, 2, 3}
	s7 := []int{1, 2, 3}
	s6Header := (*reflect.SliceHeader)(unsafe.Pointer(&s6))
	fmt.Printf("s6Header: %+v\n", s6Header)
	s7Header := (*reflect.SliceHeader)(unsafe.Pointer(&s7))
	fmt.Printf("s7Header: %+v\n", s7Header)

	fmt.Println("*****************************************")
	arr1 := [5]int{1, 2, 3, 4, 5}
	slc1 := arr1[:3]
	slc2 := append(slc1, 3)
	fmt.Println(arr1)
	fmt.Println(slc1)
	s8Header := (*reflect.SliceHeader)(unsafe.Pointer(&slc1))
	fmt.Printf("s8Header: %+v\n", s8Header)
	fmt.Println(slc2)
	s9Header := (*reflect.SliceHeader)(unsafe.Pointer(&slc2))
	fmt.Printf("s9Header: %+v\n", s9Header)
}
