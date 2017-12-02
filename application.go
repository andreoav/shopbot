package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewApplication creates a new instance of the web app
func NewApplication() *Application {
	return &Application{}
}

// Application core struct
type Application struct {
	router *mux.Router
	server *http.Server
	port   string
}

// BootstrapAndStart register routes and starts the app
func (app *Application) BootstrapAndStart() {
	app.Bootstrap()
	app.Start()
}

// Start starts the application
func (app *Application) Start() {
	app.server.ListenAndServe()
}

// Bootstrap register routes
func (app *Application) Bootstrap() {
	for path, handler := range routes {
		app.router.HandleFunc(path, handler)
	}

	app.server = &http.Server{
		Handler: app.router,
		Addr:    ":8080",
	}
}
