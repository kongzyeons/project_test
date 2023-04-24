package services

import "go_ecommerce/models"

type UserService interface {
	CreateUser(userCreate models.UserCreate) error
	LoginUser(userLogin models.UserLogin) (user_id string, err error)
	GetUser(user_id string) (result models.User_db, err error)
	GetUserOrder(user_id string) (orderHis []models.Order, err error)
}
