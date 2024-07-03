package config

import (
	"fiber-e-commerce-system-API/auth"
	"fiber-e-commerce-system-API/domain/cart"
	"fiber-e-commerce-system-API/domain/products"
	"fiber-e-commerce-system-API/domain/transaction"
	"fiber-e-commerce-system-API/domain/user"
	"fiber-e-commerce-system-API/handler"
	"fiber-e-commerce-system-API/payment"
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

	InitMidtrans()

	userRepository := user.NewRepository(DB)
	productRepository := products.NewRepository(DB)
	cartRepo := cart.NewRepository(DB)
	transactionRepo := transaction.NewRepository(DB)

	userService := user.NewService(userRepository)
	productsService := products.NewServiceProduct(productRepository)
	authService := auth.NewService()
	cartService := cart.NewService(cartRepo)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepo, paymentService)

	cartHandler := handler.NewCartHandler(cartService)
	userHandler := handler.NewUserHandler(userService, authService)
	productHandler := handler.NewProducthandler(productsService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	api := app.Group("api/v1")

	api.Get("/users", userHandler.FindAll)
	api.Post("/users/register", userHandler.RegisterUser)
	api.Post("/users/login", userHandler.Login)
	api.Get("/users/checkemail", userHandler.CheckEmailAvailable)

	api.Post("/products", productHandler.CreateProduct)
	api.Get("/products", productHandler.GetAllUser)

	api.Post("/carts", auth.AuthMiddleware(authService, userService), cartHandler.AddItemToCart)
	api.Get("/carts/:user_id", cartHandler.GetUserCart)

	api.Post("/users/transaction", auth.AuthMiddleware(authService, userService), transactionHandler.CreateTransaction)
}
