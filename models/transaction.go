package models

import "time"

type Transaction struct {
	ID              string    `bson:"_id,omitempty" json:"id"`
	Title           string    `bson:"title,omitempty" json:"title" binding:"required"`
	Amount          float64   `bson:"amount,omitempty" json:"amount" binding:"required"`
	Category        string    `bson:"category,omitempty" json:"category"`
	Date            time.Time `bson:"date,omitempty" json:"date"`
	IsExpense       bool      `bson:"is_expense,omitempty" json:"is_expense"`
	Currency        string    `bson:"currency,omitempty" json:"currency"`
	ExchangeRate    float64   `bson:"exchange_rate,omitempty" json:"exchange_rate"`
	ConvertedAmount float64   `bson:"converted_amount,omitempty" json:"converted_amount"`
}
