package handler

// LAYER 3: HANDLER (HTTP Controller / Fiber Context)

import (
	"safrenz-go-boilerplate/internal/service"
	"safrenz-go-boilerplate/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type HelloHandler struct {
	service service.HelloService
}

func NewHelloHandler(service service.HelloService) *HelloHandler {
	return &HelloHandler{service: service}
}

func (h *HelloHandler) GetHello(c *fiber.Ctx) error {
	result := h.service.SayHello()

	return utils.SuccessResponse(
		c,
		fiber.StatusOK,
		"SFR-01-001",
		"Success",
		result,
	)
}
