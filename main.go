package main

import (
	"github.com/gorilla/mux"
)

func main() {

	app := &Application{port: ":8080", router: mux.NewRouter()}

	app.BootstrapAndStart()
}
