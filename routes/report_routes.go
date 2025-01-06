package routes

import (
	"github.com/gin-gonic/gin"

	"personal-finance-api/controllers"
)

func ReportRoutes(router *gin.Engine) {
	reportRoutes := router.Group("/reports")
	{
		reportRoutes.GET("/monthly", controllers.MonthlyReport)
	}
}
