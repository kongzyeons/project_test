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

type orderHandler struct {
	orderSrv services.OrderService
}

func NewOrderHandler(orderSrv services.OrderService) OrderHandler {
	return orderHandler{orderSrv: orderSrv}
}

func (h orderHandler) CreateOrder_api(ctx *fiber.Ctx) error {
	c := config.NewFiberCtx(ctx)
	user_id := c.Params("user_id")
	orderBuy := models.OrderCreate{}
	if err := c.BodyParser(&orderBuy); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&orderBuy); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	err := h.orderSrv.CreateOrder(user_id, orderBuy)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	return c.JSON(http.StatusCreated, MessageResponse{
		Status:  true,
		Message: "Create order",
		Data:    &fiber.Map{}})
}

func (h orderHandler) DeleteOrder_api(ctx *fiber.Ctx) error {
	c := config.NewFiberCtx(ctx)
	user_id := c.Params("user_id")
	orderCancel := models.OrderDelete{}
	if err := c.BodyParser(&orderCancel); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&orderCancel); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	err := h.orderSrv.DeleteOrder(user_id, orderCancel.OrderID)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	return c.JSON(http.StatusCreated, MessageResponse{
		Status:  true,
		Message: "Delete order",
		Data:    &fiber.Map{}})
}

func (h orderHandler) GetOrder_api(ctx *fiber.Ctx) error {
	c := config.NewFiberCtx(ctx)
	user_id := c.Params("user_id")
	orderGetStatus := models.OrderGetStatus{}
	if err := c.BodyParser(&orderGetStatus); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&orderGetStatus); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}

	result, err := h.orderSrv.GetOrder(user_id, orderGetStatus.Status)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	sumPrice := 0
	for _, v := range result {
		sumPrice += int(v.SumPrice)
	}
	return c.JSON(http.StatusCreated, MessageResponse{
		Status:  true,
		Message: "Get order",
		Data:    &fiber.Map{"data": result, "total": len(result), "SumPrice": sumPrice}})
}

func (h orderHandler) UpdateOrder_api(ctx *fiber.Ctx) error {
	c := config.NewFiberCtx(ctx)
	user_id := c.Params("user_id")
	orderUpdate := models.OrderUpdate{}
	if err := c.BodyParser(&orderUpdate); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&orderUpdate); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	err := h.orderSrv.UpdateOrder(user_id, orderUpdate.OrderID, orderUpdate.Status)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	return c.JSON(http.StatusCreated, MessageResponse{
		Status:  true,
		Message: "Update order",
		Data:    &fiber.Map{}})
}
