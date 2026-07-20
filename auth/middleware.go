package auth

import (
	"fiber-e-commerce-system-API/domain/user"
	"fiber-e-commerce-system-API/helper"
	"math"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(AuthService Service, UserService user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := strings.Fields(c.Get("Authorization"))
		if len(authHeader) != 2 || !strings.EqualFold(authHeader[0], "Bearer") || authHeader[1] == "" {
			return unauthorized(c)
		}

		token, err := AuthService.ValidateToken(authHeader[1])
		if err != nil || token == nil || !token.Valid {
			return unauthorized(c)
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return unauthorized(c)
		}

		claimUserID, ok := claim["user_id"].(float64)
		if !ok || claimUserID <= 0 || claimUserID != math.Trunc(claimUserID) {
			return unauthorized(c)
		}
		userID := int(claimUserID)

		newUser, err := UserService.GetUserByID(userID)
		if err != nil || newUser.ID == 0 {
			return unauthorized(c)
		}
		c.Locals("currentUser", newUser)
		return c.Next()
	}
}

func unauthorized(c *fiber.Ctx) error {
	response := helper.APIResponse("authentication failed", http.StatusUnauthorized, "error", nil)
	return c.Status(http.StatusUnauthorized).JSON(response)
}
