package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/prongbang/cram-api/challenge"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	service := challenge.NewService()
	handle := challenge.NewHandler(service)

	v1 := app.Group("v1")
	v1.Post("/challenge/registration", handle.Registration)
	v1.Post("/challenge/request", handle.Request)
	v1.Post("/challenge/verify", handle.Verify)

	log.Fatal(app.Listen(":3000"))
}
