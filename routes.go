package main

import (
	"net/http"

	"github.com/andreoav/shopbot/handlers"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

var r = &handlers.Handlers{}

var routes = map[string]handlerFunc{
	"/": r.ProductsHandler.HomePage,
	"/products/{productId}": r.ProductsHandler.List,
}
