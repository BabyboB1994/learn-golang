package main

import (
	"fmt"
)

type Employee struct {
	ID         int
	Name, Role string
}

type Worker struct {
	ID       int
	Name     string
	Position string
}

func main() {
	var emply1 Employee

	emply1.ID = 1
	emply1.Name = "Bob"
	emply1.Role = "Developer"

	fmt.Println(emply1)
	fmt.Println("Employee ID   :", emply1.ID)
	fmt.Println("Employee Name :", emply1.Name)
	fmt.Println("Employee Role :", emply1.Role)

	var emplyOfTheMonth *Employee = &emply1
	emplyOfTheMonth.Role = "SDE-2"
	fmt.Println(*&emplyOfTheMonth.Role)
	fmt.Println((*emplyOfTheMonth).Role)
	fmt.Println(emplyOfTheMonth.Role)
	fmt.Println(emply1)
	fmt.Println()

	emply2 := emply1 // this is a copy by value
	fmt.Println(emply2)
	emply2.ID = 2
	emply2.Name = "Bryan"
	emply2.Role = "Product Manager"
	fmt.Println(emply1)

	emply3 := Employee{1, "Bob", "SDE-2"}
	fmt.Println(emply1 == emply3)

	// wkr := Worker{1, "Henry", "Manager"}
	// fmt.Println(reflect.TypeOf(Employee(wkr))) // cannot convert wkr (variable of type Worker) to Employee

}
