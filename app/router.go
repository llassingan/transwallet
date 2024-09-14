package app

import (
	"transwallet/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, walletController controller.WalletController) {

	// Group routes
	api := app.Group("/api/wallets")

	// Define routes
	api.Post("/topup", walletController.TopUp)
	api.Post("/send", walletController.SendMoney)
	api.Get("/:accountNumber", walletController.GetBalance)
	api.Get("/:accountNumber/history", walletController.GetTransactionHistory)
}
