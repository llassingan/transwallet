package controller

import (
	"github.com/gofiber/fiber/v2"
)

type WalletController interface {
	// Top-up
	TopUp(c *fiber.Ctx) error

	// Send Money
	SendMoney(c *fiber.Ctx) error

	// Get Balance
	GetBalance(c *fiber.Ctx) error

	// Get Transaction History
	GetTransactionHistory(c *fiber.Ctx) error
}
