package handler

import (
	"go_ecommerce/config"
	"go_ecommerce/models"
	"go_ecommerce/services"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) UserHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) CreateUser_api(ctx *fiber.Ctx) error {
	c := config.NewFiberCtx(ctx)
	userCreate := models.UserCreate{}
	if err := c.BodyParser(&userCreate); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&userCreate); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	err := h.userSrv.CreateUser(userCreate)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	return c.JSON(http.StatusCreated, MessageResponse{
		Status:  true,
		Message: "Create user",
		Data:    &fiber.Map{}})
}
func (h userHandler) LoginUser_api(ctx *fiber.Ctx) error {
	c := config.NewFiberCtx(ctx)
	userLogin := models.UserLogin{}
	if err := c.BodyParser(&userLogin); err != nil {
		log.Panicln("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&userLogin); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	result, err := h.userSrv.LoginUser(userLogin)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	return c.JSON(http.StatusCreated, MessageResponse{
		Status:  true,
		Message: "Login user",
		Data:    &fiber.Map{"user_id": result}})
}
func (h userHandler) GetUser_api(ctx *fiber.Ctx) error {
	c := config.NewFiberCtx(ctx)
	user_id := c.Params("user_id")
	result, err := h.userSrv.GetUser(user_id)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	return c.JSON(http.StatusCreated, MessageResponse{
		Status:  true,
		Message: "Get user",
		Data:    &fiber.Map{"data": result}})
}
func (h userHandler) GetUserOrder_api(ctx *fiber.Ctx) error {
	c := config.NewFiberCtx(ctx)
	user_id := c.Params("user_id")
	result, err := h.userSrv.GetUserOrder(user_id)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	return c.JSON(http.StatusCreated, MessageResponse{
		Status:  true,
		Message: "Get user order",
		Data:    &fiber.Map{"data": result}})
}
