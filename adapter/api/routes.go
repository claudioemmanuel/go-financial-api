package api

import (
	"financial-api/application/services"
	_ "financial-api/docs"
	"financial-api/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"financial-api/adapter/api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	userService *services.UserService,
	accountService *services.AccountService,
) {

	// Controllers
	userController := controllers.NewUserController(userService)
	accountController := controllers.NewAccountController(accountService)

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// API
	api := r.Group("/api")
	{
		// Health check
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// Login
		api.POST("/login", userController.Login)

		// Protected routes
		protected := api.Group("/protected")
		protected.Use(middleware.AuthRequired())
		{
			// Users
			protected.GET("/users", userController.GetAll)
			protected.POST("/users", userController.Create)
			protected.PUT("/users/:id", userController.Update)
			protected.DELETE("/users/:id", userController.Delete)

			// Accounts
			protected.GET("/accounts", accountController.GetAll)
			protected.POST("/accounts", accountController.Create)
			protected.PUT("/accounts/:id", accountController.Update)
			protected.DELETE("/accounts/:id", accountController.Delete)
		}
	}

	// Swagger
	swaggerConfig := &ginSwagger.Config{
		URL: "/swagger/doc.json",
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
		*c = *swaggerConfig
	}))
}
