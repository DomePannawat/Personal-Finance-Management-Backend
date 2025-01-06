package models

type Budget struct {
	ID       string  `bson:"_id,omitempty" json:"id"`
	Category string  `bson:"category,omitempty" json:"category" binding:"required"`
	Limit    float64 `bson:"limit,omitempty" json:"limit" binding:"required"`
	Spent    float64 `bson:"spent,omitempty" json:"spent"`
}
