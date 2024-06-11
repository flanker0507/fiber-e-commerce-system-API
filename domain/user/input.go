package user

type UserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Street   string `json:"street" binding:"required"`
	City     string `json:"city" binding:"required"`
	State    string `json:"state" binding:"required"`
	ZipCode  string `json:"zip_code" binding:"required"`
	Country  string `json:"country" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required"`
}
