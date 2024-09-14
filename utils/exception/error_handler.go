package exception

import (
	"errors"
	"strings"
	"transwallet/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ValidationErrorResponse defines the structure for validation error details
type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ErrorHandler handles different types of errors
func ErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error { // Check if the error is a validation error
		if ve, ok := err.(validator.ValidationErrors); ok {

			var errors []ValidationErrorResponse
			
			for _, e := range ve { // Translate field and error message
				errors = append(errors, ValidationErrorResponse{
					Field: e.Field(),
					Message:validationErrorMessage(e)})
			} // Return a detailed validation error response
			
			stdResponse := web.StdErrorResponse{
				Code: fiber.StatusBadRequest,
				Status: "Bad Request",
				Error: errors,
			}
			return c.Status(fiber.StatusBadRequest).JSON(stdResponse)
		}
		// Handle not found errors
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stdResponse := web.StdErrorResponse{
				Code: fiber.StatusNotFound,
				Status: "Not Found",
				Error: "data not found",
			}
			return c.Status(fiber.StatusNotFound).JSON(stdResponse)
		}

		if strings.Contains(err.Error(), "insufficient funds") {
			stdResponse := web.StdErrorResponse{
				Code: fiber.StatusBadRequest,
				Status: "Bad Request",
				Error: "insufficient funds",
			}
			return c.Status(fiber.StatusNotFound).JSON(stdResponse)
		}

		// For any other internal server error
		stdResponse := web.StdErrorResponse{
			Code: fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Error: "internal server error",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(stdResponse)
	}
}

// Helper function to generate validation error messages
func validationErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "minaccountid":
		return "The value is not valid"
	default:
		return "Invalid value"
	}
}
