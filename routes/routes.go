package routes

import (
	"myapp/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Tag routes
	r.POST("/tags", controllers.CreateTag)
	r.GET("/tags", controllers.ListTags)

	// Media routes
	r.POST("/media", controllers.CreateMedia)
	r.GET("/media", controllers.SearchMedia)
}
