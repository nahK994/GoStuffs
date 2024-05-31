package main

import (
	"product-management/input"
	"product-management/product"
	"strconv"
)

func main() {
	id, _ := strconv.ParseInt(input.Get("Input product Id: "), 10, 32)
	title := input.Get("Input product title: ")
	price, _ := strconv.ParseFloat(input.Get("Input product price: "), 64)
	description := input.Get("Input product description: ")
	p := product.Create(int(id), title, price, description)
	p.Print()
}
