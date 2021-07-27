package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/*
type PerishableProduct struct {
	Prod   Product
	Expiry string
}
*/

type PerishableProduct struct {
	Product
	Expiry string
}

func main() {
	//grapes := PerishableProduct{Product{200, "Grapes", 60, 30, "Food"}, "2 Days"}

	grapes := NewPerishableProduct(200, "Grapes", 60, 30, "Food", "2 Days")
	fmt.Println(grapes)
	//fmt.Println("Cost of grapes => ", grapes.Prod.Cost)

	fmt.Println("Cost of grapes => ", grapes.Product.Cost)
	fmt.Println("Cost of grapes => ", grapes.Cost)
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{Product{id, name, cost, units, category}, expiry}
}
