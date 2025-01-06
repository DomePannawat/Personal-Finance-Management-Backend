package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"personal-finance-api/config"
	"personal-finance-api/models"
)

func GetBudgets(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var budgets []models.Budget
	cursor, err := config.DB.Collection("budgets").Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching budgets"})
		return
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &budgets); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding budgets"})
		return
	}

	c.JSON(http.StatusOK, budgets)
}

func CreateBudget(c *gin.Context) {
	var budget models.Budget
	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := config.DB.Collection("budgets").InsertOne(ctx, budget)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating budget"})
		return
	}

	c.JSON(http.StatusCreated, budget)
}
