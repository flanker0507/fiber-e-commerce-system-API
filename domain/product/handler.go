package product

import (
	"errors"
	"fiber-e-commerce-system-API/domain/models"
	"fiber-e-commerce-system-API/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ProductHandler struct {
	service productService
}

func NewProducthandler(service productService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAllProduct(c *fiber.Ctx) error {
	products, err := h.service.GetAllProduct()
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			formattedErrors := helper.FormatValidateError(validationErrors)
			errorMessage := fiber.Map{"errors": formattedErrors}

			response := helper.APIResponse("Get All product failed", http.StatusUnprocessableEntity, "error", errorMessage)
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
		//Handle other type of errors
		errorMessage := fiber.Map{"errors": err.Error()}
		response := helper.APIResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Get All product success", http.StatusOK, "success", products)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var productInput ProductInput

	err := c.BodyParser(&productInput)
	if err != nil {
		var validateErrors validator.ValidationErrors
		if errors.As(err, &validateErrors) {
			formattedErrors := helper.FormatValidateError(validateErrors)
			errorMessage := fiber.Map{"errors": formattedErrors}

			response := helper.APIResponse("Create product failed", http.StatusUnprocessableEntity, "error", errorMessage)
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
		//Handle other type of errors
		errorMessage := fiber.Map{"errors": err.Error()}
		response := helper.APIResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	newProduct, err := h.service.CreateProduct(productInput)
	if err != nil {
		response := helper.APIResponse("Create product failed", http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	formatter := FormatProduct(models.Product{
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		Qty:         newProduct.Qty,
	})
	response := helper.APIResponse("Create product success", http.StatusOK, "success", formatter)
	return c.Status(http.StatusOK).JSON(response)

}
