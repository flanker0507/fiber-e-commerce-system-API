package handler

import (
	"fiber-e-commerce-system-API/domain/products"
	"fiber-e-commerce-system-API/helper"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type productHandler struct {
	productService products.Service
}

func NewProducthandler(productService products.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	var input products.ProductInput

	err := c.BodyParser(&input)
	if err != nil {
		return clientError(c, http.StatusBadRequest, "invalid request")
	}

	// Validasi menggunakan validator
	if err := validate.Struct(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": helper.FormatValidateError(err),
		})
	}

	newProduct, err := h.productService.CreateProduct(input)
	if err != nil {
		return clientError(c, http.StatusInternalServerError, "internal server error")
	}

	formatter := products.FormatProduct(newProduct)
	respone := helper.APIResponse("Create product Success", http.StatusCreated, "success", formatter)
	return c.Status(http.StatusCreated).JSON(respone)
}

func (h *productHandler) GetAllUser(c *fiber.Ctx) error {
	product, err := h.productService.GetAllProduct()
	if err != nil {
		return clientError(c, http.StatusInternalServerError, "internal server error")
	}
	respone := helper.APIResponse("Get All product Succes", http.StatusOK, "success", product)
	return c.Status(http.StatusOK).JSON(respone)
}
