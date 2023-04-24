package services

import (
	"go_ecommerce/models"
	"go_ecommerce/repository"
)

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProdcutService {
	return productService{productRepo: productRepo}
}

func (s productService) CreateProduct(productCreate models.ProductCreate) error {
	prodcutDB := models.Product_db{
		ProductName: productCreate.ProductName,
		Category:    productCreate.Category,
		Price:       productCreate.Price,
	}
	err := s.productRepo.Create(prodcutDB)
	if err != nil {
		return err
	}
	return nil
}

func (s productService) GetProduct(productID string) (result models.Product_db, err error) {
	result, err = s.productRepo.GetByID(productID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s productService) GetAllProduct() (result []models.Product_db, err error) {
	result, err = s.productRepo.GetAll()
	if err != nil {
		return result, err
	}
	return result, nil
}
