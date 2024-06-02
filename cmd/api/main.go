package main

import (
	"fmt"
	"net/http"

	"github.com/mbocsi/goapi-demo/internal/handlers"
)

func main() {
	a := &handlers.App{
		ApiHandler: &handlers.ApiHandler{
			UsersHandler:    new(handlers.UsersHandler),
			ListingsHandler: new(handlers.ListingsHandler),
		},
	}
	fmt.Println("Starting GO API service on port on localhost:8080...")
	http.ListenAndServe(":8080", a)
}
