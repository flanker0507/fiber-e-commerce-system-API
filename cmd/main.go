package main

import (
	"fiber-e-commerce-system-API/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	config.InitDB(app) //Not enough arguments in call to 'config.InitDB'

	log.Fatal(app.Listen(":8080"))
}
