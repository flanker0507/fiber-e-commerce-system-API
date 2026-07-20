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
	"net"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(app *fiber.App, cfg Config) error {
	databaseConfig := mysql.Config{
		User:      cfg.DBUser,
		Passwd:    cfg.DBPassword,
		Net:       "tcp",
		Addr:      net.JoinHostPort(cfg.DBHost, cfg.DBPort),
		DBName:    cfg.DBName,
		ParseTime: true,
		Loc:       time.Local,
	}
	dsn := databaseConfig.FormatDSN()

	var err error
	DB, err = gorm.Open(gormmysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("connect to database: %w", err)
	}

	midtransClient, err := NewMidtransClient(cfg)
	if err != nil {
		return err
	}

	userRepository := user.NewRepository(DB)
	productRepository := products.NewRepository(DB)
	cartRepo := cart.NewRepository(DB)
	transactionRepo := transaction.NewRepository(DB)

	userService := user.NewService(userRepository)
	productsService := products.NewServiceProduct(productRepository)
	authService := auth.NewService(cfg.JWTSecret)
	cartService := cart.NewService(cartRepo)
	paymentService := payment.NewService(midtransClient)
	transactionService := transaction.NewService(transactionRepo, paymentService)

	cartHandler := handler.NewCartHandler(cartService)
	userHandler := handler.NewUserHandler(userService, authService)
	productHandler := handler.NewProducthandler(productsService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	api := app.Group("api/v1")

	api.Get("/users", auth.AuthMiddleware(authService, userService), userHandler.FindAll)
	api.Post("/users/register", userHandler.RegisterUser)
	api.Post("/users/login", userHandler.Login)
	api.Get("/users/checkemail", userHandler.CheckEmailAvailable)

	api.Post("/products", productHandler.CreateProduct)
	api.Get("/products", productHandler.GetAllUser)

	api.Post("/carts", auth.AuthMiddleware(authService, userService), cartHandler.AddItemToCart)
	api.Get("/carts/:user_id", cartHandler.GetUserCart)

	api.Post("/users/transaction", auth.AuthMiddleware(authService, userService), transactionHandler.CreateTransaction)

	return nil
}
