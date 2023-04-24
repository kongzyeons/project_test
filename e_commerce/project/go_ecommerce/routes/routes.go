package routes

import (
	"go_ecommerce/config"
	"go_ecommerce/handler"
	"go_ecommerce/repository"
	"go_ecommerce/services"

	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouters(route *config.Fiber_fw, concetionDB *mongo.Client) {
	userRepository := repository.NewUserRepository(concetionDB)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	route.Post("/api/v1/create/user", userHandler.CreateUser_api)
	route.Post("/api/v1/login/user", userHandler.LoginUser_api)
	route.Get("/api/v1/getuser/:user_id", userHandler.GetUser_api)
	route.Get("/api/v1/getuser_order/:user_id", userHandler.GetUserOrder_api)

	productRepository := repository.NewProductRepository(concetionDB)
	productService := services.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	route.Post("/api/v1/create/product", productHandler.Create_api)
	route.Get("api/v1/getproduct/:product_id", productHandler.GetByID_api)
	route.Get("api/v1/getallproduct", productHandler.GetAll_api)

	orderRepository := repository.NewOrderRepository(concetionDB)
	orderService := services.NewOrderService(userRepository, productRepository, orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)
	route.Put("/api/v1/create/order/:user_id", orderHandler.CreateOrder_api)
	route.Get("/api/v1/get/order/:user_id", orderHandler.GetOrder_api)
	route.Delete("/api/v1/delete/order/:user_id", orderHandler.DeleteOrder_api)
	route.Put("/api/v1/update/order/:user_id", orderHandler.UpdateOrder_api)

}
