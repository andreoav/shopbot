package handlers

import (
	"fmt"
	"net/http"

	"github.com/andreoav/shopbot/data"
)

// ProductsHandler controller
type ProductsHandler struct {
	DefaultHandler
}

func (p *ProductsHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from golang with gorilla mux")
}

// The first thing we'll do
// We'll receive a productId and returns historical data about the prices

// List returns historical data about the given product
func (p *ProductsHandler) List(w http.ResponseWriter, r *http.Request) {
	products := data.Products{
		data.Product{Name: "Coca Cola", Price: 5.50},
		data.Product{Name: "Pepsi", Price: 5.00},
		data.Product{Name: "Fanta", Price: 5.25},
		data.Product{Name: "Fanta Guaraná", Price: 5.25},
	}

	p.asJSON(&w, products)
}
