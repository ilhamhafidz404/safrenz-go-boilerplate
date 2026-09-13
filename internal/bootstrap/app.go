package bootstrap

import (
	"safrenz-go-boilerplate/internal/handler"
	"safrenz-go-boilerplate/internal/repository"
	"safrenz-go-boilerplate/internal/service"

	"gorm.io/gorm"
)

// Inisialisasi handler, service, dan repository

type Handlers struct {
	HelloHandler *handler.HelloHandler
	UserHandler  *handler.UserHandler
}

func InitApp(db *gorm.DB) *Handlers {
	// 1. Repositories
	helloRepo := repository.NewHelloRepository()
	userRepo := repository.NewUserRepository(db)

	// 2. Services
	helloService := service.NewHelloService(helloRepo)
	userService := service.NewUserService(userRepo)

	// 3. Handlers
	helloHandler := handler.NewHelloHandler(helloService)
	userHandler := handler.NewUserHandler(userService)

	return &Handlers{
		HelloHandler: helloHandler,
		UserHandler:  userHandler,
	}
}
