package repository

import "go_ecommerce/models"

type UserRepository interface {
	Create(userCreate models.User_db) error
	Login(userLogin models.UserLogin) (result models.User_db, err error)
	GetByID(user_id string) (result models.User_db, err error)
	// AddData(user_id string, dataOrder models.Order) error
	// Update(user_id string, dataRequst models.CheckDataRequest) error
	// Delete(deleteData models.CheckDataRequest) error
}
