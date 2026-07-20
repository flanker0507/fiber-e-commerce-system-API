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
		return clientError(c, http.StatusBadRequest, "invalid request")
	}

	currentUser := c.Locals("currentUser").(models.User)

	cartItem, err := h.service.AddItemToCart(currentUser.ID, input.ProductID, input.Quantity)
	if err != nil {
		return clientError(c, http.StatusInternalServerError, "failed to process request")
	}
	response := helper.APIResponse("Item added to cart successfully", http.StatusCreated, "success", cartItem)
	return c.Status(http.StatusCreated).JSON(response)
}

func (h *cartHandler) GetUserCart(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		return clientError(c, http.StatusBadRequest, "invalid request")
	}

	cart, err := h.service.GetUserCart(userID)
	if err != nil {
		return clientError(c, http.StatusInternalServerError, "internal server error")
	}

	response := helper.APIResponse("User cart retrieved successfully", http.StatusOK, "success", cart)
	return c.Status(http.StatusOK).JSON(response)
}
