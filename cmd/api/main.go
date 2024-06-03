package main

import (
	"fmt"
	"net/http"

	"github.com/mbocsi/goapi-demo/internal/database"
	"github.com/mbocsi/goapi-demo/internal/handlers"
	"github.com/mbocsi/goapi-demo/internal/repository"
	"github.com/mbocsi/goapi-demo/internal/service"
)

func main() {
	listingRepo := repository.NewListingRepository(database.DB.Listings)
	userRepo := repository.NewUserRepository(database.DB.Users)
	listingService := service.NewListingService(listingRepo)
	userService := service.NewUserService(userRepo)
	usersHandler := handlers.NewUsersHandler(userService)
	listingsHandler := handlers.NewListingsHandler(listingService)
	apiHandler := handlers.NewApiHandler(listingsHandler, usersHandler)
	app := handlers.NewApp(apiHandler)

	fmt.Println("Starting GO API service on port on localhost:8080...")
	err := http.ListenAndServe(":8080", app)
	if err != nil {
		fmt.Println(err.Error())
	}
}
