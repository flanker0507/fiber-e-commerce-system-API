package handler

import (
	"fiber-e-commerce-system-API/auth"
	"fiber-e-commerce-system-API/domain/user"
	"fiber-e-commerce-system-API/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type userHandler struct {
	service user.Service
	auth    auth.Service
}

func NewUserHandler(service user.Service, auth auth.Service) *userHandler {
	return &userHandler{service, auth}
}

var validate = validator.New()

func (h *userHandler) RegisterUser(c *fiber.Ctx) error {
	var input user.RegisterUserInput
	var validate = validator.New()

	err := c.BodyParser(&input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}

	// Validasi menggunakan validator
	if err := validate.Struct(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": helper.FormatValidateError(err),
		})
	}

	registerUser, err := h.service.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}

	token, err := h.auth.GenerateToken(registerUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)

	}

	formatter := user.FormatUser(registerUser, token)
	response := helper.APIResponse("Account has been registered", http.StatusCreated, "success", formatter)
	return c.JSON(response)
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	var input user.LoginInput
	var validate = validator.New()

	err := c.BodyParser(&input)
	if err != nil {
		response := helper.APIResponse("Login account failed", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}

	// Validasi menggunakan validator
	if err := validate.Struct(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": helper.FormatValidateError(err),
		})
	}

	loginUser, err := h.service.Login(input)
	if err != nil {
		response := helper.APIResponse("Login account failed", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}
	token, err := h.auth.GenerateToken(loginUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}

	formatter := user.FormatUser(loginUser, token)
	response := helper.APIResponse("Login Success", http.StatusCreated, "success", formatter)
	return c.JSON(response)
}

func (h *userHandler) FindAll(c *fiber.Ctx) error {
	findAll, err := h.service.GetAllUser()
	if err != nil {
		response := helper.APIResponse("Get All account failed", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}

	response := helper.APIResponse("Login Success", http.StatusOK, "success", findAll)
	return c.JSON(response)
}

func (h *userHandler) CheckEmailAvailable(c *fiber.Ctx) error {
	var input user.CheckEmailInput

	err := c.BodyParser(&input)
	if err != nil {
		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}
	// Validasi menggunakan validator
	if err := validate.Struct(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": helper.FormatValidateError(err),
		})
	}

	isEmailAvailable, err := h.service.IsEmailAvailable(input)
	if err != nil {
		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}

	data := fiber.Map{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	return c.JSON(response)
}
