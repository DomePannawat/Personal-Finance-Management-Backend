package main

import (
	"log"

	"personal-finance-api/config"
	"personal-finance-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	config.ConnectDB()

	router := gin.Default()

	// router
	routes.TransactionRoutes(router)
	routes.BudgetRoutes(router)

	routes.AuthRoutes(router)
	routes.CategoryRoutes(router)
	routes.ReportRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Personal Finance API")
	})

	// Start server
	port := config.GetEnv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
