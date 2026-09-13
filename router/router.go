package router

// Mendaftarkan Route API & Menghubungkan Handler + Middleware

import (
	"safrenz-go-boilerplate/internal/handler"

	"github.com/gofiber/fiber/v2"
)

// all handler router need
type RouterConfig struct {
	App          *fiber.App
	HelloHandler *handler.HelloHandler
	UserHandler  *handler.UserHandler
}

func SetupRouter(config *RouterConfig) {
	api := config.App.Group("/api/v1")

	api.Get("/hello", config.HelloHandler.GetHello)
	api.Get("/users", config.UserHandler.GetUsers)
}
