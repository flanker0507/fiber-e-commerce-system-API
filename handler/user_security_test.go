package handler

import (
	"errors"
	"fiber-e-commerce-system-API/auth"
	"fiber-e-commerce-system-API/domain/models"
	"fiber-e-commerce-system-API/domain/user"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type securityTestUserService struct {
	users       []models.User
	findErr     error
	loginErr    error
	currentUser models.User
	currentErr  error
}

func (s securityTestUserService) RegisterUser(user.RegisterUserInput) (models.User, error) {
	return models.User{}, nil
}

func (s securityTestUserService) Login(user.LoginInput) (models.User, error) {
	return models.User{}, s.loginErr
}

func (s securityTestUserService) IsEmailAvailable(user.CheckEmailInput) (bool, error) {
	return false, nil
}

func (s securityTestUserService) GetUserByID(int) (models.User, error) {
	return s.currentUser, s.currentErr
}

func (s securityTestUserService) GetAllUser() ([]models.User, error) {
	return s.users, s.findErr
}

func (s securityTestUserService) UpdateUser(user.FormUpdateUserInput) (models.User, error) {
	return models.User{}, nil
}

func TestFindAllUsersReturnsSanitizedDTOs(t *testing.T) {
	service := securityTestUserService{users: []models.User{{
		ID:       1,
		Name:     "Safe User",
		Email:    "safe@example.com",
		Password: "password-hash-marker",
		Role:     "admin",
		Token:    "token-marker",
	}}, currentUser: models.User{ID: 1, Name: "Authenticated User"}}
	authService := auth.NewService("unit-test-jwt-secret")
	handler := NewUserHandler(service, authService)
	app := fiber.New()
	app.Get("/users", auth.AuthMiddleware(authService, service), handler.FindAll)

	token, err := authService.GenerateToken(service.currentUser.ID)
	if err != nil {
		t.Fatalf("GenerateToken() returned an error: %v", err)
	}
	request := httptest.NewRequest("GET", "/users", nil)
	request.Header.Set("Authorization", "Bearer "+token)
	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() returned an error: %v", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("io.ReadAll() returned an error: %v", err)
	}

	for _, forbidden := range []string{"password-hash-marker", "token-marker", `"Password"`, `"password"`, `"Token"`, `"token"`, `"Role"`, `"role"`} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("user response contains forbidden content %q", forbidden)
		}
	}
	for _, required := range []string{`"id"`, `"name"`, `"email"`, `"created_at"`, `"updated_at"`} {
		if !strings.Contains(string(body), required) {
			t.Fatalf("user response is missing safe field %q", required)
		}
	}
}

func TestFindAllUsersRequiresValidJWT(t *testing.T) {
	service := securityTestUserService{
		users:       []models.User{{ID: 1, Password: "password-hash-marker", Token: "token-marker"}},
		currentUser: models.User{ID: 1},
	}
	authService := auth.NewService("unit-test-jwt-secret")
	handler := NewUserHandler(service, authService)
	app := fiber.New()
	app.Get("/users", auth.AuthMiddleware(authService, service), handler.FindAll)
	missingClaimToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{}).SignedString([]byte("unit-test-jwt-secret"))
	if err != nil {
		t.Fatalf("SignedString() returned an error: %v", err)
	}

	tests := []struct {
		name          string
		authorization string
	}{
		{name: "missing token"},
		{name: "invalid token", authorization: "Bearer invalid-token-marker"},
		{name: "invalid scheme", authorization: "Basic invalid-token-marker"},
		{name: "missing user claim", authorization: "Bearer " + missingClaimToken},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", "/users", nil)
			if test.authorization != "" {
				request.Header.Set("Authorization", test.authorization)
			}
			response, err := app.Test(request)
			if err != nil {
				t.Fatalf("app.Test() returned an error: %v", err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				t.Fatalf("io.ReadAll() returned an error: %v", err)
			}

			if response.StatusCode != fiber.StatusUnauthorized {
				t.Fatalf("response status = %d; want %d", response.StatusCode, fiber.StatusUnauthorized)
			}
			for _, forbidden := range []string{"password-hash-marker", "token-marker", "invalid-token-marker"} {
				if strings.Contains(string(body), forbidden) {
					t.Fatalf("unauthorized response contains forbidden content %q", forbidden)
				}
			}
		})
	}
}

func TestLoginDoesNotExposeRawServiceError(t *testing.T) {
	const internalDetail = "database driver internal detail marker"
	service := securityTestUserService{loginErr: errors.New(internalDetail)}
	handler := NewUserHandler(service, nil)
	app := fiber.New()
	app.Post("/login", handler.Login)

	request := httptest.NewRequest("POST", "/login", strings.NewReader(`{"email":"safe@example.com","password":"valid-input"}`))
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() returned an error: %v", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("io.ReadAll() returned an error: %v", err)
	}

	if response.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("response status = %d; want %d", response.StatusCode, fiber.StatusUnauthorized)
	}
	if strings.Contains(string(body), internalDetail) {
		t.Fatal("login response exposes the raw service error")
	}
	if !strings.Contains(string(body), "authentication failed") {
		t.Fatal("login response does not contain the sanitized error message")
	}
}

func TestGlobalErrorHandlerDoesNotExposeRawInternalError(t *testing.T) {
	const internalDetail = "filesystem and stack trace detail marker"
	app := fiber.New(fiber.Config{ErrorHandler: ErrorHandler})
	app.Get("/failure", func(*fiber.Ctx) error {
		return errors.New(internalDetail)
	})

	response, err := app.Test(httptest.NewRequest("GET", "/failure", nil))
	if err != nil {
		t.Fatalf("app.Test() returned an error: %v", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("io.ReadAll() returned an error: %v", err)
	}

	if response.StatusCode != fiber.StatusInternalServerError {
		t.Fatalf("response status = %d; want %d", response.StatusCode, fiber.StatusInternalServerError)
	}
	if strings.Contains(string(body), internalDetail) {
		t.Fatal("global error response exposes the raw internal error")
	}
	if !strings.Contains(string(body), "internal server error") {
		t.Fatal("global error response does not contain the sanitized error message")
	}
}
