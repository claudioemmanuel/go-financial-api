package api

import (
	_ "financial-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"financial-api/adapter/api/controllers"
	"financial-api/application/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, userService *services.UserService) {
	userController := controllers.NewUserController(userService)

	// API
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		api.GET("/users", userController.GetAll)
		api.POST("/users", userController.Create)
		api.PUT("/users/:id", userController.Update)
		api.DELETE("/users/:id", userController.Delete)
	}

	// Serve swagger.json
	// r.GET("/swagger/json", func(c *gin.Context) {
	// 	c.File("./docs/swagger.json")
	// })

	// Swagger
	swaggerConfig := &ginSwagger.Config{
		URL: "/swagger/doc.json",
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
		*c = *swaggerConfig
	}))
}
