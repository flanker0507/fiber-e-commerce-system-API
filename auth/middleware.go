package auth

import (
	"fiber-e-commerce-system-API/domain/user"
	"fiber-e-commerce-system-API/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func AuthMiddleware(AuthService Service, UserService user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := AuthService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		userID := int(claim["user_id"].(float64))

		newUser, err := UserService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			return c.Status(http.StatusUnauthorized).JSON(response)
		}
		c.Locals("currentUser", newUser) // Use c.Locals to set the user object
		return c.Next()
	}
}
