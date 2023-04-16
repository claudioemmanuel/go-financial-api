package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"financial-api/adapter/api"
	"financial-api/application/services"
	"financial-api/config/database"
	g "financial-api/infrastructure/persistence/gorm"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Could not initialize the database: ", err)
	}

	// Dependency injection
	userRepository := g.NewUserRepositoryGorm(db)
	accountRepository := g.NewAccountRepositoryGorm(db)

	// Create the services
	userService := services.NewUserService(userRepository)
	accountService := services.NewAccountService(accountRepository)

	// Set up the Gin web framework
	r := gin.Default()

	// Register routes pass the array of services
	api.RegisterRoutes(r, userService, accountService)

	// Start the server
	r.Run()
}
