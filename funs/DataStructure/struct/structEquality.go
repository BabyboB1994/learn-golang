package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	ID   int
	Name string
	Role string
}

var manager struct {
	ID   int
	Role string
	Name string
}

var worker struct {
	ID   int
	Name string
	Role string
}

type student struct{
	Id int
	Name *string
}

type meter int

func main() {
	var rad1 meter = 2
	var rad2 int = 3
	fmt.Println(reflect.TypeOf(rad1))
	fmt.Println(reflect.TypeOf(rad2))
	// fmt.Println(rad1 == rad2) // cannot compare rad1 == rad2 (mismatched types meter and int)

	emp := Employee{1, "Glenn", "IT Manager"}
	worker = struct {
		ID   int
		Name string
		Role string
	}{2, "Steve", "CEO"}
	fmt.Println(worker == emp) // why this is not erroring out even though worker and emp are of different types??
	// If all the fields of a struct are comparable, the struct itself is comparable, so two expressions of that type may be compared using == or !=

	manager = struct {
		ID   int
		Role string
		Name string
	}{1, "IT Manager", "Glenn"}
	fmt.Println(manager)
	// fmt.Println(emp == manager) // cannot compare emp == manager (mismatched types Employee and struct{ID int; Role string; Name string})

	// checking reflect.DeepEqual()
	std1Name := "Bob"
	std2Name := "Bob"

	student1 := student{1, &std1Name}
	student2 := student{1, &std2Name,}
	fmt.Printf("%#v\n",student1)
	fmt.Printf("%#v\n",student2)
	fmt.Println(student1 == student2) //false
	fmt.Println(reflect.DeepEqual(student1, student2)) //true - reflect.DeepEqual() checks the checks which the fields are pointing to
}
