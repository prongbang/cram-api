package challenge

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/prongbang/cram-api/core"
)

type Handler interface {
	Registration(c *fiber.Ctx) error
	Request(c *fiber.Ctx) error
	Verify(c *fiber.Ctx) error
}

type handler struct {
	Service Service
}

// Registration implements Handler
func (h *handler) Registration(c *fiber.Ctx) error {
	r := Registration{}
	if err := c.BodyParser(&r); err != nil {
		return core.BadRequest(c, nil)
	}

	status := h.Service.Registration(r)
	if status {
		return core.Ok(c, nil)
	}

	return core.BadRequest(c, nil)
}

// Request implements Handler
func (h *handler) Request(c *fiber.Ctx) error {
	r := Request{}
	if err := c.BodyParser(&r); err != nil {
		return core.BadRequest(c, nil)
	}

	challengeText := h.Service.Request(r)

	return core.Ok(c, Challenge{
		Challenge: challengeText,
	})
}

// Verify implements Handler
func (h *handler) Verify(c *fiber.Ctx) error {
	v := Verify{}
	if err := c.BodyParser(&v); err != nil {
		return core.BadRequest(c, nil)
	}

	verified := h.Service.Verify(v)
	if verified {
		return core.Ok(c, core.Token{
			Token: uuid.NewString(),
		})
	}

	return core.BadRequest(c, "Verify Error")
}

func NewHandler(service Service) Handler {
	return &handler{
		Service: service,
	}
}
