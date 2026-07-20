package main

import (
	"fiber-e-commerce-system-API/config"
	"fiber-e-commerce-system-API/handler"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("application configuration is invalid")
	}

	app := fiber.New(fiber.Config{ErrorHandler: handler.ErrorHandler})
	if err := config.InitDB(app, cfg); err != nil {
		log.Fatal("application initialization failed")
	}

	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatal("server stopped unexpectedly")
	}
}
