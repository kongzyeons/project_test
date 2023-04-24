package handler

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	CreateUser_api(ctx *fiber.Ctx) error
	LoginUser_api(ctx *fiber.Ctx) error
	GetUser_api(ctx *fiber.Ctx) error
	GetUserOrder_api(ctx *fiber.Ctx) error
}
