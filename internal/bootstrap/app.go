package bootstrap

import (
	"safrenz-go-boilerplate/internal/handler"
	"safrenz-go-boilerplate/internal/repository"
	"safrenz-go-boilerplate/internal/service"
	"safrenz-go-boilerplate/pkg/platform"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Handlers struct {
	HelloHandler *handler.HelloHandler
	UserHandler  *handler.UserHandler
}

func InitApp(db *gorm.DB, rdb *redis.Client) *Handlers {
	roleClient := platform.NewRoleClient(rdb)

	// 2. Repositories
	helloRepo := repository.NewHelloRepository()
	userRepo := repository.NewUserRepository(db)

	// 3. Services
	helloService := service.NewHelloService(helloRepo)
	userService := service.NewUserService(userRepo, roleClient)

	// 4. Handlers
	helloHandler := handler.NewHelloHandler(helloService)
	userHandler := handler.NewUserHandler(userService)

	return &Handlers{
		HelloHandler: helloHandler,
		UserHandler:  userHandler,
	}
}
