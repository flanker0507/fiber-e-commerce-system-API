package user

import (
	"fiber-e-commerce-system-API/domain/models"
	"time"
)

// UserResponse is the public representation of a user. It deliberately omits
// credentials and authorization data stored on the database model.
type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuthResponse is returned only after successful authentication. The token is
// intentionally included here rather than exposed through the database model.
type AuthResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func FormatUser(user models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FormatUsers(users []models.User) []UserResponse {
	responses := make([]UserResponse, 0, len(users))
	for _, model := range users {
		responses = append(responses, FormatUser(model))
	}
	return responses
}

func FormatAuthResponse(user models.User, token string) AuthResponse {
	return AuthResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}
}
