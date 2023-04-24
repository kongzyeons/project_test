package models

type ProductCreate struct {
	ProductName string  `json:"product_name" bson:"product_name" validate:"required"`
	Category    string  `json:"category" bson:"category" validate:"required"`
	Price       float64 `json:"price" bson:"price" validate:"required"`
}
