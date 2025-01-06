package routes

import (
	"github.com/gin-gonic/gin"

	"personal-finance-api/controllers"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")
	{
		categoryRoutes.GET("/", controllers.GetCategories)
		categoryRoutes.POST("/", controllers.CreateCategory)
		categoryRoutes.DELETE("/:id", controllers.DeleteCategory)
	}
}
