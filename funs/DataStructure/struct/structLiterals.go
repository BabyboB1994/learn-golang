package main

import "fmt"

type book struct {
	isbn      string
	name      string
	pageCount int
}

func main() {
	book1 := book{"ST101", "Story of panther", 200} // order-based form of struct literals. Field(s) cannot be omitted.
	book2 := book{isbn: "NV203", pageCount: 300}    // field:value form of struct literals. Field(s) can be omitted.

	fmt.Println(book1)
	fmt.Println(book2.name)
}
