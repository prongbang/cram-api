package core

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Token struct {
	Token string `json:"token"`
}

type Response struct {
	Code    int `json:"code"`
	Data    any `json:"data,omitempty"`
	Message any `json:"message"`
}

func Ok(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusOK).JSON(&Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	})
}

func BadRequest(c *fiber.Ctx, data any) error {
	response := &Response{
		Code:    http.StatusBadRequest,
		Message: http.StatusText(http.StatusBadRequest),
	}
	if data != nil {
		response.Message = data
	}
	return c.Status(http.StatusBadRequest).JSON(response)
}
