package handler

import "github.com/gofiber/fiber/v2"

type ProductHandler interface {
	Create_api(ctx *fiber.Ctx) error
	GetByID_api(ctx *fiber.Ctx) error
	GetAll_api(ctx *fiber.Ctx) error
}
