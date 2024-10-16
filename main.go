package main

import (
	"myapp/controllers"
	"myapp/routes"
	"github.com/gin-gonic/gin"
	_ "myapp/docs"
	"github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// @title Media and Tag API
// @version 1.0
// @description REST API for creating and searching tags and media items

func main() {
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080
}
