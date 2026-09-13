package main

import (
	"log"

	"safrenz-go-boilerplate/config"
	"safrenz-go-boilerplate/database/seeders"
	"safrenz-go-boilerplate/internal/bootstrap"
	"safrenz-go-boilerplate/router"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadEnv()

	// Database, automigrate & Seeder
	db := config.InitDB()
	seeders.RunAll(db)

	// Inisialisasi Seluruh Handler via Bootstrap
	handlers := bootstrap.InitApp(db)

	app := fiber.New(fiber.Config{
		AppName: "Safrenz Go Boilerplate v1.0",
	})

	// Setup Router
	router.SetupRouter(&router.RouterConfig{
		App:          app,
		HelloHandler: handlers.HelloHandler,
		UserHandler:  handlers.UserHandler,
	})

	// Menjalankan Server
	log.Println("Server aktif di http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
