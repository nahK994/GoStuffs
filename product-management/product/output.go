package product

import (
	"fmt"
	"os"
)

func (product *Product) Print() {
	fmt.Println(product.id, product.title, product.price, product.description)
	file, _ := os.Create(product.title + ".txt")
	info := fmt.Sprintf("Id: %v\nTitle: %v\nPrice: %v\nDescription: %v", product.id, product.title, product.price, product.description)
	file.WriteString(info)
	file.Close()
}
