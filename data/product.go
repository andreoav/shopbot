package data

type Product struct {
	Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Products []Product
