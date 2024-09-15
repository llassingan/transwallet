package core

import (
	"os"
	"path/filepath"
	"transwallet/api/controller"

	// "transwallet/utils"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	// "github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App, walletController controller.WalletController) {

	// group routes
	api := app.Group("/api/wallets")

	// define routes
	api.Post("/topup", walletController.TopUp)
	api.Post("/send", walletController.SendMoney)
	api.Get("/:accountNumber", walletController.GetBalance)
	api.Get("/:accountNumber/history", walletController.GetTransactionHistory)

	docs := app.Group("/docs/wallets")
	docs.Get("/swagger.json", func(c *fiber.Ctx) error {
		// Get the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Unable to get current directory")
		}

		// Construct the path to the apispec.json file
		jsonPath := filepath.Join(cwd, "docs", "apispec.json")

		// Check if the file exists
		if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).SendString("Swagger JSON file not found")
		}

		// Set the appropriate content type
		c.Set("Content-Type", "application/json")

		// Send the file
		return c.SendFile(jsonPath)
	})

	// Serve Swagger UI
	docs.Get("/swagger/*", fiberSwagger.FiberWrapHandler(fiberSwagger.URL("/docs/wallets/swagger.json"))) // The URL pointing to API definitio)
}
