package main

import (
	"fmt"
	"net/http"

	"github.com/mbocsi/goapi-demo/internal/database"
	"github.com/mbocsi/goapi-demo/internal/handlers"
	"github.com/mbocsi/goapi-demo/internal/middleware"
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
	corsWrapper := middleware.NewResponseHeader(app, "Access-Control-Allow-Origin", "*")

	fmt.Println("Starting GO API service on http://localhost:8080...")
	fmt.Println(`   _____  ____             _____ _____   _____  ______ __  __  ____  
  / ____|/ __ \      /\   |  __ \_   _| |  __ \|  ____|  \/  |/ __ \ 
 | |  __| |  | |    /  \  | |__) || |   | |  | | |__  | \  / | |  | |
 | | |_ | |  | |   / /\ \ |  ___/ | |   | |  | |  __| | |\/| | |  | |
 | |__| | |__| |  / ____ \| |    _| |_  | |__| | |____| |  | | |__| |
  \_____|\____/  /_/    \_\_|   |_____| |_____/|______|_|  |_|\____/ 
                                                                     
                                                                     `)
	err := http.ListenAndServe(":8080", corsWrapper)
	if err != nil {
		fmt.Println(err.Error())
	}
}
