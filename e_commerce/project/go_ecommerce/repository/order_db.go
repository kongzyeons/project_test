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

type orderRepository struct {
	db *mongo.Client
}

func NewOrderRepository(db *mongo.Client) OrderRepository {
	return orderRepository{db: db}
}
func (r orderRepository) getCollection() *mongo.Collection {
	collection := r.db.Database("Ecommerce").Collection("User")
	return collection
}

func (r orderRepository) AddData(user_id string, orderBuy models.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := models.User_db{}

	// Check user_id
	quary := bson.M{"user_id": user_id}
	err := r.getCollection().FindOne(ctx, quary).Decode(&user)
	if err != nil {
		err := fmt.Errorf("not found user_id")
		return err
	}

	// // Check order id
	order_id := uuid.New().String()
	quary_order := bson.M{
		"order.order_id": order_id,
	}
	err = r.getCollection().FindOne(ctx, quary_order).Decode(&user)
	if err == nil {
		order_id := uuid.New().String()
		quary_order := bson.M{
			"order.order_id": order_id,
		}
		err = r.getCollection().FindOne(ctx, quary_order).Decode(&user)
	}
	orderBuy.OrderID = order_id
	status_log := models.StatusOrder{
		Status:     "waitting",
		StatusTime: time.Now(),
	}
	orderBuy.StatusLog = append(orderBuy.StatusLog, status_log)
	quary = bson.M{"user_id": user_id}
	update := bson.M{
		"$push": bson.M{"order": orderBuy},
	}
	_, err = r.getCollection().UpdateOne(ctx, quary, update)
	if err != nil {
		return err
	}
	return nil
}

func (r orderRepository) DeleteData(user_id, order_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := models.User_db{}

	// check user_id
	quary_user := bson.M{"user_id": user_id}
	err := r.getCollection().FindOne(ctx, quary_user).Decode(&user)
	if err != nil {
		err := fmt.Errorf("not found user_id")
		return err
	}

	// check order_id
	quary_order := bson.M{
		"order": bson.M{
			"$elemMatch": bson.M{
				"order_id": order_id,
				"status_log.status": bson.M{
					"$nin": []string{"settle", "reject", "success"},
				},
			},
		},
	}

	err = r.getCollection().FindOne(ctx, quary_order).Decode(&user)
	if err != nil {
		err := fmt.Errorf("not found order_id")
		return err
	}
	quary_order = bson.M{
		"order": bson.M{
			"order_id": order_id,
		},
	}
	remove := bson.M{
		"$pull": quary_order,
	}
	_, err = r.getCollection().UpdateOne(ctx, quary_user, remove)
	if err != nil {
		return err
	}
	return nil
}

func (r orderRepository) UpdateData(user_id, order_id, status string, status_list []string, lenStatus int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := models.User_db{}

	// check user_id
	quary_user := bson.M{"user_id": user_id}
	err := r.getCollection().FindOne(ctx, quary_user).Decode(&user)
	if err != nil {
		err := fmt.Errorf("not found user_id")
		return err
	}
	// check order_id
	// Example "$nin": []string{"settle", "reject", "success"} >> waiting
	quary_order := bson.M{
		"user_id": user_id,
		"order": bson.M{
			"$elemMatch": bson.M{
				"order_id": order_id,
				"status_log": bson.M{
					"$size": lenStatus,
				},
				"status_log.status": bson.M{
					"$nin": status_list,
				},
			},
		},
	}
	err = r.getCollection().FindOne(ctx, quary_order).Decode(&user)

	if err != nil {
		err := fmt.Errorf("not found order_id")
		return err
	}

	status_log := models.StatusOrder{
		Status:     status,
		StatusTime: time.Now(),
	}

	user.OrderHis[0].StatusLog = append(user.OrderHis[0].StatusLog, status_log)

	quary := bson.M{"user_id": user_id, "order.order_id": order_id}

	update := bson.M{
		"$push": bson.M{
			"order.$.status_log": status_log,
		},
	}

	_, err = r.getCollection().UpdateOne(ctx, quary, update)
	if err != nil {
		return err
	}

	return nil
}

// func (r orderRepository) GetData(user_id, status string, status_list []string, lenStatus int) (result models.User_db, err error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	// check user_id
// 	quary_user := bson.M{"user_id": user_id}
// 	err = r.getCollection().FindOne(ctx, quary_user).Decode(&result)
// 	if err != nil {
// 		err := fmt.Errorf("not found user_id")
// 		return result, err
// 	}
// 	quary_order := bson.M{
// 		"user_id": user_id,
// 		"order": bson.M{
// 			"$elemMatch": bson.M{
// 				// "status_log": bson.M{
// 				// 	"$size": 3,
// 				// },
// 				"status_log.status": status,
// 			},
// 		},
// 	}

// 	err = r.getCollection().FindOne(ctx, quary_order).Decode(&result)
// 	if err != nil {
// 		err := fmt.Errorf("not found order_id")
// 		return result, err
// 	}
// 	return result, nil
// }

// func (r orderRepository) GetData(user_id string) (result models.User_db, err error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	quary := bson.M{"user_id": user_id}
// 	err = r.getCollection().FindOne(ctx, quary).Decode(&result)
// 	if err != nil {
// 		return result, err
// 	}
// 	quary_order := bson.M{
// 		"order.status_log.status": bson.M{
// 			"$nin": []string{"reject", "settle", "success"},
// 		},
// 	}
// 	cursor, err := r.getCollection().Find(ctx, quary_order)
// 	if err != nil {
// 		return result, err
// 	}
// 	if err = cursor.All(context.TODO(), &result); err != nil {
// 		return result, err
// 	}
// 	return result, nil
// }
