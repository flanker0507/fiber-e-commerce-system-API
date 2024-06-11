package user

import (
	"errors"
	"fiber-e-commerce-system-API/domain/models"
	"fiber-e-commerce-system-API/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler struct {
	userService *UserService
}

func NewHandler(userService *UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var userInput UserInput

	err := c.BodyParser(&userInput)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			formattedErrors := helper.FormatValidateError(validationErrors)
			errorMessage := fiber.Map{"errors": formattedErrors}

			response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
			return c.Status(http.StatusUnprocessableEntity).JSON(response)
		}

		// Handle other types of errors
		errorMessage := fiber.Map{"errors": err.Error()}
		response := helper.APIResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	newUser, err := h.userService.CreateUser(userInput)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", nil)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	formatter := FormatUser(newUser)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) Index(c *fiber.Ctx) error {
	users, err := h.userService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": nil})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"users": users})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var input LoginInput

	err := c.BodyParser(&input)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			formattedErrors := helper.FormatValidateError(validationErrors)
			errorMessage := fiber.Map{"errors": formattedErrors}

			response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
			return c.Status(http.StatusUnprocessableEntity).JSON(response)
		}

		// Handle other types of errors
		errorMessage := fiber.Map{"errors": err.Error()}
		response := helper.APIResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := fiber.Map{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	formatter := FormatUser(loggedinUser)
	response := helper.APIResponse("Successfully login", http.StatusOK, "success", formatter)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		response := helper.APIResponse("Invalid User ID", http.StatusNotFound, "error", nil)
		return c.Status(http.StatusNotFound).JSON(response)
	}

	var userInput models.User
	if err := c.BodyParser(&userInput); err != nil {
		var validationError validator.ValidationErrors
		if errors.As(err, &validationError) {
			formattedErrors := helper.FormatValidateError(validationError)
			errorMessage := fiber.Map{"errors": formattedErrors}

			response := helper.APIResponse("Update account failed", http.StatusUnprocessableEntity, "error", errorMessage)
			return c.Status(http.StatusUnprocessableEntity).JSON(response)
		}

		errorMessage := fiber.Map{"errors": err.Error()}
		response := helper.APIResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	updatedUser, err := h.userService.UpdateUser(uint(userID), userInput)
	if err != nil {
		response := helper.APIResponse("Failed to update user", http.StatusInternalServerError, "error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("User successfully updated", http.StatusOK, "success", updatedUser)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		response := helper.APIResponse("Invalid User ID", http.StatusNotFound, "eroor", nil)
		return c.Status(http.StatusNotFound).JSON(response)
	}

	deleterUser, err := h.userService.DeleteUser(uint(userID))
	if err != nil {
		response := helper.APIResponse("Failed to deletd User Id", http.StatusInternalServerError, "error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("User successfully deleted", http.StatusOK, "success", deleterUser)
	return c.Status(http.StatusOK).JSON(response)

}
