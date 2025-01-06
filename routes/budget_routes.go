package routes

import (
	"personal-finance-api/controllers"

	"github.com/gin-gonic/gin"
)

func BudgetRoutes(router *gin.Engine) {
	budgetRoutes := router.Group("/budgets")
	{
		budgetRoutes.GET("/", controllers.GetBudgets)
		budgetRoutes.POST("/", controllers.CreateBudget)
	}
}
