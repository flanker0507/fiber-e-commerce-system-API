package config

import (
	"fiber-e-commerce-system-API/domain/models"
	"fiber-e-commerce-system-API/domain/user"
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

	DB.AutoMigrate(&models.User{}, &models.Cart{}, &models.Product{}, &models.Payment{}, &models.CartProduct{})

	userRepository := user.NewRespository(DB)

	userService := user.NewUserService(userRepository)

	userHandler := user.NewHandler(userService)

	app.Post("/users/register", userHandler.CreateUser)
	app.Post("/users", userHandler.Login)
	app.Get("/users", userHandler.Index)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)

}
