package handler

import (
	"github.com/gofiber/fiber/v2"
)

type MessageResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

func err_response(err error) MessageResponse {
	return MessageResponse{
		Status:  false,
		Message: "error",
		Data:    &fiber.Map{"data": err.Error()}}
}
