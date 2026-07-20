package handler

import (
	"errors"
	"fiber-e-commerce-system-API/helper"

	"github.com/gofiber/fiber/v2"
)

func clientError(c *fiber.Ctx, status int, message string) error {
	response := helper.APIResponse(message, status, "error", nil)
	return c.Status(status).JSON(response)
}

// ErrorHandler is the last-resort Fiber error handler. It prevents unexpected
// internal errors from being returned verbatim when a route does not handle
// them itself.
func ErrorHandler(c *fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	message := "internal server error"

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		status = fiberError.Code
		switch status {
		case fiber.StatusBadRequest:
			message = "invalid request"
		case fiber.StatusUnauthorized:
			message = "authentication failed"
		case fiber.StatusForbidden:
			message = "access forbidden"
		case fiber.StatusNotFound:
			message = "resource not found"
		case fiber.StatusMethodNotAllowed:
			message = "method not allowed"
		default:
			if status >= 400 && status < 500 {
				message = "failed to process request"
			}
		}
	}

	return clientError(c, status, message)
}
