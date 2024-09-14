package controller

import (
	"transwallet/service"
	"transwallet/model/web"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type WalletControllerImpl struct {
	WalletService service.WalletService
}

// Create a new instance of WalletController
func NewWalletController(walletService service.WalletService) *WalletControllerImpl {
	return &WalletControllerImpl{WalletService: walletService}
}

// Handle Top-Up request
func (ctrl *WalletControllerImpl) TopUp(c *fiber.Ctx) error {
	var req web.TopUpRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	
	response,err := ctrl.WalletService.TopUp(c.Context(), req)
	if err != nil {
		return err
	}
	stdResponse := web.StdResponse{
		Code: fiber.StatusOK,
		Status: "Success",
		Data: response,
	}
	return c.Status(fiber.StatusOK).JSON(stdResponse)
}

// Handle Send Money request
func (ctrl *WalletControllerImpl) SendMoney(c *fiber.Ctx) error {
	var req web.SendRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	response,err := ctrl.WalletService.SendMoney(c.Context(), req)
	if err != nil {
		return err
	}
	stdResponse := web.StdResponse{
		Code: fiber.StatusOK,
		Status: "Success",
		Data: response,
	}
	return c.Status(fiber.StatusOK).JSON(stdResponse)
}

// Handle Get Balance request
func (ctrl *WalletControllerImpl) GetBalance(c *fiber.Ctx) error {
	accountIdStr := c.Params("accountNumber")
	accountId, err := strconv.ParseUint(accountIdStr, 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid account Number")
	}

	response, err := ctrl.WalletService.GetBalance(c.Context(), uint(accountId))
	if err != nil {
		return err
	}
	stdResponse := web.StdResponse{
		Code: fiber.StatusOK,
		Status: "Success",
		Data: response,
	}

	return c.Status(fiber.StatusOK).JSON(stdResponse)
}

// Handle Get Transaction History request
func (ctrl *WalletControllerImpl) GetTransactionHistory(c *fiber.Ctx) error {
	accountIdStr := c.Params("accountNumber")
	accountId, err := strconv.ParseUint(accountIdStr, 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid account Number")
	}

	response, err := ctrl.WalletService.GetTransactionHistory(c.Context(), uint(accountId))
	if err != nil {
		return err
	}
	stdResponse := web.StdResponse{
		Code: fiber.StatusOK,
		Status: "Success",
		Data: response,
	}
	return c.Status(fiber.StatusOK).JSON(stdResponse)
}
