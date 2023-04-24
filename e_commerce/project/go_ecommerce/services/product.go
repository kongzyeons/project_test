package services

import "go_ecommerce/models"

type ProdcutService interface {
	CreateProduct(productCreate models.ProductCreate) error
	GetProduct(productID string) (result models.Product_db, err error)
	GetAllProduct() (result []models.Product_db, err error)
}
