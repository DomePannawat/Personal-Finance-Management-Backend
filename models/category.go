package models

type Category struct {
	ID          string `bson:"_id,omitempty" json:"id"`
	Name        string `bson:"name" json:"name" binding:"required"`
	Description string `bson:"description,omitempty" json:"description"`
}
