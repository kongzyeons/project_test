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

type prodcutRepository struct {
	db *mongo.Client
}

func NewProductRepository(db *mongo.Client) ProductRepository {
	return prodcutRepository{db: db}
}

func (r prodcutRepository) getCollection() *mongo.Collection {
	collection := r.db.Database("Ecommerce").Collection("Product")
	return collection
}

func (r prodcutRepository) Create(productCreate models.Product_db) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check product_name or Category
	quary := bson.M{
		"$and": []bson.M{
			{"product_name": productCreate.ProductName},
			{"category": productCreate.Category},
		},
	}
	err := r.getCollection().FindOne(ctx, quary).Decode(&productCreate)
	if err == nil {
		err := fmt.Errorf("product_name or category not ready")
		return err
	}

	// Check product id
	product_id := uuid.New().String()
	err = r.getCollection().FindOne(ctx, bson.M{"product_id": product_id}).Decode(&productCreate)
	productCreate.ProductID = product_id
	for err == nil {
		product_id := uuid.New().String()
		err = r.getCollection().FindOne(ctx, bson.M{"product_id": product_id}).Decode(&productCreate)
		productCreate.ProductID = product_id
	}
	newProduct := models.Product_db{
		ProductID:   productCreate.ProductID,
		ProductName: productCreate.ProductName,
		Category:    productCreate.Category,
		Price:       productCreate.Price,
		CreateDate:  time.Now(),
		UpdateDate:  time.Now(),
	}
	_, err = r.getCollection().InsertOne(ctx, newProduct)
	if err != nil {
		return err
	}
	return nil

}

func (r prodcutRepository) GetByID(productID string) (result models.Product_db, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	quary := bson.M{"product_id": productID}
	err = r.getCollection().FindOne(ctx, quary).Decode(&result)
	if err != nil {
		err := fmt.Errorf("not found product_id")
		return result, err
	}
	return result, nil
}

func (r prodcutRepository) GetAll() (result []models.Product_db, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := r.getCollection().Find(ctx, bson.M{})
	if err != nil {
		err := fmt.Errorf("not found product")
		return result, err
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		err := fmt.Errorf("not found product")
		return nil, err
	}
	return result, err
}
