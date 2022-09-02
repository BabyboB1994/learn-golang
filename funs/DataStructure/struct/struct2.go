package main

import (
	"fmt"
)

type Employee struct {
	ID         int
	Name, Role string
}

var employeeArr = []Employee{
	{ID: 1, Name: "Bob", Role: "SDE-2"},
	{ID: 2, Name: "Bryan", Role: "Product Manager"},
	{ID: 3, Name: "Kelly", Role: "MD"}}

func main() {
	employee := getEmployeeById(4, employeeArr)
	fmt.Println(employee.ID)
	fmt.Println(employee.Name == "")
	employee2 := Employee{}
	fmt.Println(employee2)

	getEmployeeById(1, employeeArr).Name = "Brian"
}

func getEmployeeById(id int, employees []Employee) Employee {
	var employee Employee
	for _, emp := range employees {
		if emp.ID == id {
			employee = emp
			break
		}
	}
	//return employee
	return Employee{23, "Tom", "Dev"}
}
