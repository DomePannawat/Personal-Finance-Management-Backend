package routes

import (
	"personal-finance-api/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(router *gin.Engine) {
	transactionRoutes := router.Group("/transactions")
	{
		transactionRoutes.GET("/", controllers.GetTransactions)
		transactionRoutes.POST("/", controllers.CreateTransaction)
	}
}
