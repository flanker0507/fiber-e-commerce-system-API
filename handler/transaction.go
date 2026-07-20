package handler

import (
	"fiber-e-commerce-system-API/domain/models"
	"fiber-e-commerce-system-API/domain/transaction"
	"fiber-e-commerce-system-API/helper"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var input transaction.CreateTransactionInput

	err := c.BodyParser(&input)
	if err != nil {
		return clientError(c, http.StatusBadRequest, "invalid request")
	}

	// Get user ID from JWT
	currentUser := c.Locals("currentUser").(models.User)
	input.UserID = currentUser.ID
	input.User = currentUser

	transaction, err := h.service.CreateTransaction(input)
	if err != nil {
		return clientError(c, http.StatusInternalServerError, "failed to process request")
	}
	response := helper.APIResponse("Transaction created successfully", http.StatusCreated, "success", transaction)
	return c.Status(http.StatusCreated).JSON(response)
}

func (h *transactionHandler) GetTransactionByID(c *fiber.Ctx) error {
	transactionID, err := c.ParamsInt("id")
	if err != nil {
		return clientError(c, http.StatusBadRequest, "invalid request")
	}

	transaction, err := h.service.GetTransactionsByUserID(transactionID)
	if err != nil {
		return clientError(c, http.StatusInternalServerError, "internal server error")
	}

	response := helper.APIResponse("Transaction retrieved successfully", http.StatusOK, "success", transaction)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *transactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		return clientError(c, http.StatusInternalServerError, "internal server error")
	}

	response := helper.APIResponse("Transactions retrieved successfully", http.StatusOK, "success", transactions)
	return c.Status(http.StatusOK).JSON(response)
}
