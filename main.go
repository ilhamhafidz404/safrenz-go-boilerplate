package main

import (
	"log"

	"safrenz-go-boilerplate/config"
	"safrenz-go-boilerplate/database/seeders"
	"safrenz-go-boilerplate/internal/bootstrap"
	"safrenz-go-boilerplate/router"

	_ "safrenz-go-boilerplate/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title           Safrenz Go Boilerplate API
// @version         1.0
// @description     API Documentation for Safrenz Go Boilerplate.
// @host            localhost:3000
// @BasePath        /
func main() {

	config.LoadEnv()

	db := config.InitDB()
	rdb := config.InitRedis()

	// Automigrate & Seeder
	seeders.RunAll(db)

	handlers := bootstrap.InitApp(db, rdb)

	app := fiber.New(fiber.Config{
		AppName: "Safrenz Go Boilerplate v1.0",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	router.SetupRouter(&router.RouterConfig{
		App:          app,
		HelloHandler: handlers.HelloHandler,
		UserHandler:  handlers.UserHandler,
	})

	log.Println("Server aktif di http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
