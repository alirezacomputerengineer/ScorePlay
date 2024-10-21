package main

import (
	"myapp/routes"
	"github.com/gin-gonic/gin"
	_ "myapp/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
