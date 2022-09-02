package main

import "fmt"

type CatalogItem struct{
	id int
	Name string
	CatalogBrand
}

type CatalogBrand struct{
	Bid int
	Brand string
}

type Product struct{
	int
	string
	CatalogBrand
}

func main(){
	item1 := CatalogItem{101, "Sneakers", CatalogBrand{Bid:90, Brand:"Vans"}}
	fmt.Printf("%#v\n", item1)
	fmt.Println("Item brand details")
	fmt.Println(item1.Bid)
	fmt.Println(item1.Brand)

	// When anonymous types are used, it takes implicit names. Hence we cannot have more than 1 anynonymous types of the same type
	product1 := Product{901, "Sports", CatalogBrand{Bid:201, Brand:"Cosco"}}
	fmt.Printf("%#v\n", product1) // main.Product{int:901, string:"Sports", CatalogBrand:main.CatalogBrand{Bid:201, Brand:"Cosco"}}
	fmt.Println(product1.int)
	fmt.Println(product1.string)
}