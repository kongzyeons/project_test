package repository

import "go_ecommerce/models"

type OrderRepository interface {
	AddData(user_id string, dataOrder models.Order) error
	DeleteData(user_id, order_id string) error
	UpdateData(user_id, order_id, status string, status_list []string, lenStatus int) error
}
