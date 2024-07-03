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
		errors := helper.FormatValidateError(err)
		errorMessage := fiber.Map{"errors": errors}

		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	// Get user ID from JWT
	currentUser := c.Locals("currentUser").(models.User)
	input.UserID = currentUser.ID

	transaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}
	response := helper.APIResponse("Transaction created successfully", http.StatusCreated, "success", transaction)
	return c.JSON(response)
}

func (h *transactionHandler) GetTransactionByID(c *fiber.Ctx) error {
	transactionID, err := c.ParamsInt("id")
	if err != nil {
		response := helper.APIResponse("Invalid transaction ID", http.StatusBadRequest, "error", err.Error())
		return c.JSON(response)
	}

	transaction, err := h.service.GetTransactionsByUserID(transactionID)
	if err != nil {
		response := helper.APIResponse("Failed to get transaction", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}

	response := helper.APIResponse("Transaction retrieved successfully", http.StatusOK, "success", transaction)
	return c.JSON(response)
}

func (h *transactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		response := helper.APIResponse("Failed to get transactions", http.StatusUnprocessableEntity, "error", err.Error())
		return c.JSON(response)
	}

	response := helper.APIResponse("Transactions retrieved successfully", http.StatusOK, "success", transactions)
	return c.JSON(response)
}
