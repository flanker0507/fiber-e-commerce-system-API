package user

type RegisterUserInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type FormUpdateUserInput struct {
	ID    int
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Error error
}
