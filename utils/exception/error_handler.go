package exception

import (
	"errors"
	"strings"
	"transwallet/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// define struct for validation error details
type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// handling error
func ErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		// handling validation error
		if ve, ok := err.(validator.ValidationErrors); ok {

			var errors []ValidationErrorResponse

			for _, e := range ve {
				// get the  field and error message
				errors = append(errors, ValidationErrorResponse{
					Field:   e.Field(),
					Message: validationErrorMessage(e)})
			}
			// wrap error standard response
			stdResponse := web.StdErrorResponse{
				Code:   fiber.StatusBadRequest,
				Status: "Bad Request",
				Error:  errors,
			}
			return c.Status(fiber.StatusBadRequest).JSON(stdResponse)
		}


		if strings.Contains(err.Error(), "sender not found") {
			stdResponse := web.StdErrorResponse{
				Code:   fiber.StatusBadRequest,
				Status: "Bad Request",
				Error:  "sender not found",
			}
			return c.Status(fiber.StatusBadRequest).JSON(stdResponse)
		}

		if strings.Contains(err.Error(), "recepient not found") {
			stdResponse := web.StdErrorResponse{
				Code:   fiber.StatusBadRequest,
				Status: "Bad Request",
				Error:  "recepient not found",
			}
			return c.Status(fiber.StatusBadRequest).JSON(stdResponse)
		}


		// handling not found errors
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stdResponse := web.StdErrorResponse{
				Code:   fiber.StatusNotFound,
				Status: "Not Found",
				Error:  "data not found",
			}
			return c.Status(fiber.StatusNotFound).JSON(stdResponse)
		}

		// handling invalid input
		if strings.Contains(err.Error(), "Invalid account Number") {
			stdResponse := web.StdErrorResponse{
				Code:   fiber.StatusBadRequest,
				Status: "Bad Request",
				Error:  "The value is not valid account number",
			}
			return c.Status(fiber.StatusBadRequest).JSON(stdResponse)
		}


		// handling insufficient funds
		if strings.Contains(err.Error(), "insufficient funds") {
			stdResponse := web.StdErrorResponse{
				Code:   fiber.StatusBadRequest,
				Status: "Bad Request",
				Error:  "insufficient funds",
			}
			return c.Status(fiber.StatusBadRequest).JSON(stdResponse)
		}

		// handling any other error
		stdResponse := web.StdErrorResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Error:  err,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(stdResponse)
	}
}

// helper function to generate validation error messages
func validationErrorMessage(e validator.FieldError) string {
	// check used tags
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "minaccountid":
		return "The value is not valid account number"
	case "numeric":
		return "The value is not valid account number"
	case "min":
		return "The minimum transaction is 10"
	default:
		return "Invalid value"
	}
}
