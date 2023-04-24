package models

type UserCreate struct {
	Username  string `json:"username" bson:"username" validate:"required"`
	Password  string `json:"password" bson:"password" validate:"required"`
	Firstname string `json:"firstname" bson:"firstname" validate:"required"`
	Lastname  string `json:"lastname" bson:"lastname" validate:"required"`
}

type UserLogin struct {
	Username string `json:"username" bson:"username" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}
