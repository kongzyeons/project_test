package repository

import (
	"context"
	"fmt"
	"go_ecommerce/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) UserRepository {
	return userRepository{db: db}
}

func (r userRepository) getCollection() *mongo.Collection {
	collection := r.db.Database("Ecommerce").Collection("User")
	return collection
}

func (r userRepository) Create(userCreate models.User_db) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check username
	quary := bson.M{"username": userCreate.Username}
	err := r.getCollection().FindOne(ctx, quary).Decode(&userCreate)
	if err == nil {
		err := fmt.Errorf("username not ready")
		return err
	}

	// Check user id
	user_id := uuid.New().String()
	err = r.getCollection().FindOne(ctx, bson.M{"user_id": user_id}).Decode(&userCreate)
	userCreate.UserID = user_id
	for err == nil {
		user_id := uuid.New().String()
		err = r.getCollection().FindOne(ctx, bson.M{"user_id": user_id}).Decode(&userCreate)
		userCreate.UserID = user_id
	}

	newUser := models.User_db{
		UserID:     userCreate.UserID,
		Username:   userCreate.Username,
		Password:   userCreate.Password,
		Firstname:  userCreate.Firstname,
		Lastname:   userCreate.Lastname,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
		OrderHis:   []models.Order{},
	}

	_, err = r.getCollection().InsertOne(ctx, newUser)
	if err != nil {
		return err
	}
	return nil
}

func (r userRepository) Login(userLogin models.UserLogin) (result models.User_db, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	quary := bson.M{
		"username": userLogin.Username,
		"password": userLogin.Password,
	}
	err = r.getCollection().FindOne(ctx, quary).Decode(&result)
	if err != nil {
		err := fmt.Errorf("Invalid username or password.")
		return result, err
	}
	return result, nil
}

func (r userRepository) GetByID(user_id string) (result models.User_db, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	quary := bson.M{"user_id": user_id}
	err = r.getCollection().FindOne(ctx, quary).Decode(&result)
	if err != nil {
		err := fmt.Errorf("not found user_id")
		return result, err
	}
	return result, nil
}

// func (r userRepository) AddData(user_id string, orderBuy models.Order) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	userCreate := models.User_db{}
// 	// // Check order id
// 	order_id := uuid.New().String()
// 	quary := bson.M{
// 		"order": bson.M{
// 			"$in": bson.M{
// 				"order_id": order_id,
// 			},
// 		},
// 	}
// 	err := r.getCollection().FindOne(ctx, quary).Decode(&userCreate)
// 	if err == nil {
// 		order_id := uuid.New().String()
// 		quary = bson.M{
// 			"order": bson.M{
// 				"$in": bson.M{
// 					"order_id": order_id,
// 				},
// 			},
// 		}
// 		err = r.getCollection().FindOne(ctx, quary).Decode(&userCreate)
// 	}

// 	orderBuy.OrderID = order_id
// 	status_log := models.StatusOrder{
// 		Status:     "waiting",
// 		StatusTime: time.Now(),
// 	}
// 	orderBuy.StatusLog = append(orderBuy.StatusLog, status_log)
// 	quary = bson.M{"user_id": user_id}
// 	update := bson.M{
// 		"$push": bson.M{"order": orderBuy},
// 	}
// 	_, err = r.getCollection().UpdateOne(ctx, quary, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r userRepository) Update(user_id string, dataRequst models.CheckDataRequest) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	selectQuary := bson.M{}
// 	data := checkDataRequest(dataRequst)
// 	for key, value := range data {
// 		if value {
// 			selectQuary[key] = "$" + key
// 		}
// 	}
// 	quary := selectQuary
// 	_, err := r.getCollection().UpdateOne(ctx, bson.M{"user_id": user_id}, bson.M{"$set": quary})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r userRepository) Delete(deleteData models.CheckDataRequest) error {
// 	return nil
// }
