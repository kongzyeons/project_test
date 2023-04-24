package services

import (
	"fmt"
	"go_ecommerce/models"
	"go_ecommerce/repository"
)

type orderService struct {
	userRepo    repository.UserRepository
	productRepo repository.ProductRepository
	orderRepo   repository.OrderRepository
}

func NewOrderService(userRepo repository.UserRepository, productRepo repository.ProductRepository, orderRepo repository.OrderRepository) OrderService {
	return orderService{
		userRepo:    userRepo,
		productRepo: productRepo,
		orderRepo:   orderRepo,
	}
}

func (s orderService) CreateOrder(user_id string, orederBuy models.OrderCreate) error {
	oreder := models.Order{}
	// Check product_id
	product, err := s.productRepo.GetByID(orederBuy.ProductID)
	if err != nil {
		return err
	}
	for i := 0; i < orederBuy.Amount; i++ {
		oreder.Product = append(oreder.Product, product)
	}

	err = s.orderRepo.AddData(user_id, oreder)
	if err != nil {
		return err
	}
	return nil
}

func (s orderService) DeleteOrder(user_id, order_id string) error {
	err := s.orderRepo.DeleteData(user_id, order_id)
	if err != nil {
		return err
	}
	return nil
}

func (s orderService) GetOrder(user_id, status string) (result []models.OrderGet, err error) {
	// check Status
	if !checkStatus(status) {
		err := fmt.Errorf("status must be {waitting,settle, reject, success}")
		return result, err
	}
	user, err := s.userRepo.GetByID(user_id)
	if err != nil {
		return result, err
	}
	var lenStatus int
	if status == "waitting" {
		lenStatus = 1
	} else if status == "settle" {
		lenStatus = 2
	} else if status == "reject" {
		lenStatus = 3
	} else if status == "success" {
		lenStatus = 3
	}
	for _, v := range user.OrderHis {
		// get status log waitting
		if len(v.StatusLog) == lenStatus && v.StatusLog[lenStatus-1].Status == status {
			order := models.OrderGet{
				OrderID:     v.OrderID,
				ProductID:   v.Product[0].ProductID,
				ProductName: v.Product[0].ProductName,
				Category:    v.Product[0].Category,
				Price:       v.Product[0].Price,
				Amount:      len(v.Product),
				SumPrice:    v.Product[0].Price * float64(len(v.Product)),
			}
			result = append(result, order)
		}
	}
	return result, nil
}

func (s orderService) UpdateOrder(user_id, order_id, status string) error {
	// check Status
	if !checkStatus(status) || status == "waitting" {
		err := fmt.Errorf("status must be {settle, reject, success}")
		return err
	}

	var lenStatus int
	status_list := []string{}
	if status == "settle" {
		lenStatus = 1
		status_list = []string{"settle", "reject", "success"}
	} else if status == "reject" {
		lenStatus = 2
		status_list = []string{"reject", "success"}
	} else if status == "success" {
		lenStatus = 2
		status_list = []string{"reject", "success"}
	}
	err := s.orderRepo.UpdateData(user_id, order_id, status, status_list, lenStatus)
	if err != nil {
		return err
	}
	return nil
}

// func (s orderService) GetOrder(user_id, status string) (result []models.OrderGet, err error) {
// 	// check Status
// 	if !checkStatus(status) {
// 		err := fmt.Errorf("status must be {waitting,settle, reject, success}")
// 		return result, err
// 	}
// 	var lenStatus int
// 	status_list := []string{}
// 	if status == "waitting" {
// 		lenStatus = 1
// 		status_list = []string{status}
// 	} else if status == "settle" {
// 		lenStatus = 2
// 		status_list = []string{status}
// 	} else if status == "reject" {
// 		lenStatus = 3
// 		status_list = []string{status}
// 	} else if status == "success" {
// 		lenStatus = 3
// 		status_list = []string{status}
// 	}

// 	user, err := s.orderRepo.GetData(user_id, status, status_list, lenStatus)
// 	if err != nil {
// 		return result, err
// 	}
// 	for _, v := range user.OrderHis {
// 		order := models.OrderGet{
// 			OrderID:     v.OrderID,
// 			ProductID:   v.Product[0].ProductID,
// 			ProductName: v.Product[0].ProductName,
// 			Category:    v.Product[0].Category,
// 			Price:       v.Product[0].Price,
// 			Amount:      len(v.Product),
// 			SumPrice:    v.Product[0].Price * float64(len(v.Product)),
// 		}
// 		result = append(result, order)
// 	}

// 	return result, nil
// }
