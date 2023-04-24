package models

import "time"

type User_db struct {
	UserID     string    `json:"user_id" bson:"user_id"`
	Username   string    `json:"username" bson:"username"`
	Password   string    `json:"password" bson:"password"`
	Firstname  string    `json:"firstname" bson:"firstname"`
	Lastname   string    `json:"lastname" bson:"lastname"`
	CreateDate time.Time `json:"create_date" bson:"create_date"`
	UpdateDate time.Time `json:"update_date" bson:"update_date"`
	OrderHis   []Order   `json:"order" bson:"order"`
}
