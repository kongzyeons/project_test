package handler

import "github.com/gofiber/fiber/v2"

type OrderHandler interface {
	CreateOrder_api(ctx *fiber.Ctx) error
	DeleteOrder_api(ctx *fiber.Ctx) error
	GetOrder_api(ctx *fiber.Ctx) error
	UpdateOrder_api(ctx *fiber.Ctx) error
}
