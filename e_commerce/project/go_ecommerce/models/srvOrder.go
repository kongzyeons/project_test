package models

type OrderCreate struct {
	ProductID string `json:"product_id" bson:"product_id" validate:"required"`
	Amount    int    `json:"amount" bson:"amount" validate:"required"`
}

type OrderDelete struct {
	OrderID string `json:"order_id" bson:"order_id" validate:"required"`
}

type OrderGet struct {
	OrderID     string  `json:"order_id" bson:"order_id" validate:"required"`
	ProductID   string  `json:"product_id" bson:"product_id" validate:"required"`
	ProductName string  `json:"product_name" bson:"product_name" validate:"required"`
	Category    string  `json:"category" bson:"category" validate:"required"`
	Price       float64 `json:"price" bson:"price" validate:"required"`
	Amount      int     `json:"amount" bson:"amount" validate:"required"`
	SumPrice    float64 `json:"sumPrice" bson:"sumPrice" validate:"required"`
}

type OrderUpdate struct {
	OrderID string `json:"order_id" bson:"order_id" validate:"required"`
	Status  string `json:"status" bson:"status" validate:"required"`
}

type OrderGetStatus struct {
	Status string `json:"status" bson:"status" validate:"required"`
}
