package models

import "time"

type AuditLog struct {
	UserID    string    `bson:"user_id" json:"user_id"`
	Action    string    `bson:"action" json:"action"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Data      string    `bson:"data" json:"data"`
}
