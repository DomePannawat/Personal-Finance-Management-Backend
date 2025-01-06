package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"personal-finance-api/config"
	"personal-finance-api/models"
)

func MonthlyReport(c *gin.Context) {
	month := c.Query("month")
	year := c.Query("year")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var transactions []models.Transaction
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "date", Value: bson.D{
				{Key: "$gte", Value: time.Date(toInt(year), time.Month(toInt(month)), 1, 0, 0, 0, 0, time.UTC)},
				{Key: "$lt", Value: time.Date(toInt(year), time.Month(toInt(month)+1), 1, 0, 0, 0, 0, time.UTC)},
			}},
		}},
	}

	cursor, err := config.DB.Collection("transactions").Aggregate(ctx, mongo.Pipeline{matchStage})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching report"})
		return
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &transactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding report"})
		return
	}

	var totalIncome, totalExpense float64
	for _, t := range transactions {
		if t.IsExpense {
			totalExpense += t.Amount
		} else {
			totalIncome += t.Amount
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"total_income":  totalIncome,
		"total_expense": totalExpense,
	})
}

func toInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}
