package repository

import "go_ecommerce/models"

type ProductRepository interface {
	Create(productCreate models.Product_db) error
	GetByID(productID string) (result models.Product_db, err error)
	GetAll() (result []models.Product_db, err error)
}
