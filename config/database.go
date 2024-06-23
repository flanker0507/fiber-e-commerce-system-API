package config

import (
	"fiber-e-commerce-system-API/auth"
	"fiber-e-commerce-system-API/domain/user"
	"fiber-e-commerce-system-API/handler"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB(app *fiber.App) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error Connecting to the database: %v", err)
	}

	userRepository := user.NewRepository(DB)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	api := app.Group("api/v1")

	api.Get("/users", userHandler.FindAll)
	api.Post("/users/register", userHandler.RegisterUser)
	api.Post("/users/login", userHandler.Login)
	api.Get("/users/checkemail", userHandler.CheckEmailAvailable)
}
