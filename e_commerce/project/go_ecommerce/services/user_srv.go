package services

import (
	"go_ecommerce/models"
	"go_ecommerce/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) CreateUser(userCreate models.UserCreate) error {
	userCreateDB := models.User_db{
		Username:  userCreate.Username,
		Password:  userCreate.Password,
		Firstname: userCreate.Firstname,
		Lastname:  userCreate.Lastname,
	}
	err := s.userRepo.Create(userCreateDB)
	if err != nil {
		return err
	}
	return nil
}
func (s userService) LoginUser(userLogin models.UserLogin) (user_id string, err error) {
	result, err := s.userRepo.Login(userLogin)
	if err != nil {
		return user_id, err
	}
	user_id = result.UserID
	return user_id, nil
}
func (s userService) GetUser(user_id string) (result models.User_db, err error) {
	result, err = s.userRepo.GetByID(user_id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (s userService) GetUserOrder(user_id string) (orderHis []models.Order, err error) {
	result, err := s.userRepo.GetByID(user_id)
	if err != nil {
		return orderHis, err
	}
	for _, v := range result.OrderHis {
		orderHis = append(orderHis, v)
	}
	return orderHis, nil

}
