package handler

import (
	"fiber-e-commerce-system-API/domain/cart"
	"fiber-e-commerce-system-API/domain/models"
	"fiber-e-commerce-system-API/helper"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type cartHandler struct {
	service cart.Service
}

func NewCartHandler(service cart.Service) *cartHandler {
	return &cartHandler{service}
}

func (h *cartHandler) AddItemToCart(c *fiber.Ctx) error {
	var input cart.AddProductInput

	err := c.BodyParser(&input)
	if err != nil {
		errors := helper.FormatValidateError(err)
		errorMessage := fiber.Map{"errors": errors}

		response := helper.APIResponse("Failed Add To Cart", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	currentUser := c.Locals("currentUser").(models.User)

	cartItem, err := h.service.AddItemToCart(currentUser.ID, input.ProductID, input.Quantity)
	if err != nil {
		response := helper.APIResponse("Failed to add item to cart", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}
	response := helper.APIResponse("Item added to cart successfully", http.StatusCreated, "success", cartItem)
	return c.JSON(response)
}

func (h *cartHandler) GetUserCart(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		response := helper.APIResponse("Invalid user ID", http.StatusBadRequest, "error", err.Error())
		return c.JSON(response)
	}

	cart, err := h.service.GetUserCart(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user cart", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}

	response := helper.APIResponse("User cart retrieved successfully", http.StatusOK, "success", cart)
	return c.JSON(response)
}
