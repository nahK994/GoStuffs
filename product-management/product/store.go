package product

type Product struct {
	id          int
	title       string
	price       float64
	description string
}

func Create(id int, title string, price float64, description string) *Product {
	product := Product{
		id:          id,
		title:       title,
		price:       price,
		description: description,
	}
	return &product
}

func (p *Product) UpdatePrice(price float64) {
	p.price = price
}
