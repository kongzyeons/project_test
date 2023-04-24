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

type productHandler struct {
	productSrv services.ProdcutService
}

func NewProductHandler(productSrv services.ProdcutService) ProductHandler {
	return productHandler{productSrv: productSrv}
}

func (h productHandler) Create_api(ctx *fiber.Ctx) error {
	c := config.NewFiberCtx(ctx)
	productCreate := models.ProductCreate{}
	if err := c.BodyParser(&productCreate); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&productCreate); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	err := h.productSrv.CreateProduct(productCreate)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	return c.JSON(http.StatusCreated, MessageResponse{
		Status:  true,
		Message: "Create product",
		Data:    &fiber.Map{}})
}

func (h productHandler) GetByID_api(ctx *fiber.Ctx) error {
	c := config.NewFiberCtx(ctx)
	product_id := c.Params("product_id")
	result, err := h.productSrv.GetProduct(product_id)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	return c.JSON(http.StatusCreated, MessageResponse{
		Status:  true,
		Message: "Get product",
		Data:    &fiber.Map{"data": result}})
}

func (h productHandler) GetAll_api(ctx *fiber.Ctx) error {
	c := config.NewFiberCtx(ctx)
	result, err := h.productSrv.GetAllProduct()
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	return c.JSON(http.StatusCreated, MessageResponse{
		Status:  true,
		Message: "Get product",
		Data:    &fiber.Map{"data": result, "total": len(result)}})
}
