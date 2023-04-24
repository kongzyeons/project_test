package services

import "go_ecommerce/models"

type OrderService interface {
	CreateOrder(user_id string, orederBuy models.OrderCreate) error
	DeleteOrder(user_id, order_id string) error
	GetOrder(user_id, status string) (result []models.OrderGet, err error)
	UpdateOrder(user_id, order_id, status string) error
}

func checkStatus(status string) bool {
	status_list := []string{"waitting", "settle", "reject", "success"}
	for _, v := range status_list {
		if status == v {
			return true
		}
	}
	return false
}
