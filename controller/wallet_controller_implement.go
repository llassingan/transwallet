package controller

import (
	// "encoding/json"
	"strconv"
	"transwallet/model/web"
	"transwallet/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type WalletControllerImpl struct {
	WalletService service.WalletService
	Logger        *logrus.Logger
}

// create a new instance of WalletController
func NewWalletController(walletService service.WalletService, logger *logrus.Logger) *WalletControllerImpl {
	return &WalletControllerImpl{WalletService: walletService, Logger: logger}
}

// handle top up request
func (ctrl *WalletControllerImpl) TopUp(c *fiber.Ctx) error {
	var req web.TopUpRequest

	if err := c.BodyParser(&req); err != nil {
		// error handling
		ctrl.Logger.Error(err)
		ctrl.Logger.WithFields(logrus.Fields{
			"accountNumber": req.AccountID,
			"amount":        req.Amount,
			"status":        fiber.StatusBadRequest,
		}).Error("Invalid input")
		return err
	}
	//logging
	ctrl.Logger.WithFields(logrus.Fields{
		"accountNumber": req.AccountID,
		"amount":        req.Amount,
		"method":        c.Method(),
		"path":          c.Path(),
	}).Info("top up request")

	// pass value to service layer
	response, err := ctrl.WalletService.TopUp(c.Context(), req)
	if err != nil {
		return err
	}
	// wrap response
	stdResponse := web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "Success",
		Data:   response,
	}
	return c.Status(fiber.StatusOK).JSON(stdResponse)
}

// handle send money request
func (ctrl *WalletControllerImpl) SendMoney(c *fiber.Ctx) error {
	var req web.SendRequest

	if err := c.BodyParser(&req); err != nil {
		ctrl.Logger.Error(err)
		ctrl.Logger.WithFields(logrus.Fields{
			"fromAccount": req.FromAccount,
			"toAccount":   req.ToAccount,
			"amount":      req.Amount,
			"status":      fiber.StatusBadRequest,
		}).Error("Invalid input")
		return err
	}

	ctrl.Logger.WithFields(logrus.Fields{
		"fromAccount": req.FromAccount,
		"toAccount":   req.ToAccount,
		"amount":      req.Amount,
		"method":      c.Method(),
		"path":        c.Path(),
	}).Info("send money request")

	response, err := ctrl.WalletService.SendMoney(c.Context(), req)
	if err != nil {
		return err
	}
	stdResponse := web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "Success",
		Data:   response,
	}
	return c.Status(fiber.StatusOK).JSON(stdResponse)
}

// handle get balance request
func (ctrl *WalletControllerImpl) GetBalance(c *fiber.Ctx) error {
	accountIdStr := c.Params("accountNumber")
	accountId, err := strconv.Atoi(accountIdStr)
	if err != nil {
		ctrl.Logger.WithFields(logrus.Fields{
			"accountId": accountId,
			"status":    fiber.StatusBadRequest,
		}).Error("Invalid account Number")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid account Number")
	}

	if accountId < 100001 || accountId > 999999 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid account Number")
	}

	ctrl.Logger.WithFields(logrus.Fields{
		"accountId": accountId,
		"method":    c.Method(),
		"path":      c.Path(),
	}).Info("get balance request")

	response, err := ctrl.WalletService.GetBalance(c.Context(), accountId)
	if err != nil {
		return err
	}
	stdResponse := web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "Success",
		Data:   response,
	}

	return c.Status(fiber.StatusOK).JSON(stdResponse)
}

// handle get transaction history request
func (ctrl *WalletControllerImpl) GetTransactionHistory(c *fiber.Ctx) error {
	accountIdStr := c.Params("accountNumber")
	accountId, err := strconv.Atoi(accountIdStr)
	if err != nil {
		ctrl.Logger.WithFields(logrus.Fields{
			"accountId": accountId,
			"status":    fiber.StatusBadRequest,
		}).Error("Invalid account Number")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid account Number")
	}

	if accountId < 100001 || accountId > 999999 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid account Number")
	}

	ctrl.Logger.WithFields(logrus.Fields{
		"accountId": accountId,
		"method":    c.Method(),
		"path":      c.Path(),
	}).Info("get history request")

	response, err := ctrl.WalletService.GetTransactionHistory(c.Context(), accountId)
	if err != nil {
		return err
	}

	if len(response) == 0 {
		spcres := "no data"
		stdResponse := web.StdResponse{
			Code:   fiber.StatusOK,
			Status: "Success",
			Data:   spcres,
		}
		return c.Status(fiber.StatusOK).JSON(stdResponse)
	}

	stdResponse := web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "Success",
		Data:   response,
	}
	return c.Status(fiber.StatusOK).JSON(stdResponse)
}
