package user

import "fiber-e-commerce-system-API/domain/models"

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Address string `json:"address"`
}

func FormatUser(user models.User) UserResponse {
	address := user.Street + ", " + user.City + user.State + ", " + user.ZipCode + user.Country
	return UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Role:    user.Role,
		Address: address,
	}
}
