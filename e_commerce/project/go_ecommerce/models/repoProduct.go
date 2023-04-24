package models

import "time"

type Product_db struct {
	ProductID   string    `json:"product_id" bson:"product_id"`
	ProductName string    `json:"product_name" bson:"product_name"`
	Category    string    `json:"category" bson:"category"`
	Price       float64   `json:"price" bson:"price"`
	CreateDate  time.Time `json:"create_date" bson:"create_date"`
	UpdateDate  time.Time `json:"update_date" bson:"update_date"`
}
