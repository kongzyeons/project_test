package models

import "time"

type Order struct {
	OrderID   string        `json:"order_id" bson:"order_id"`
	Product   []Product_db  `json:"product" bson:"product"`
	StatusLog []StatusOrder `json:"status_log" bson:"status_log"`
}

type StatusOrder struct {
	Status     string    `json:"status" bson:"status"`
	StatusTime time.Time `json:"status_time" bson:"status_time"`
}
