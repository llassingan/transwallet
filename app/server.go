package app

import (
	"transwallet/controller"
	"transwallet/db"
	"transwallet/repository"
	"transwallet/service"
	"transwallet/utils/exception"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewServer() *fiber.App {
	option := []validator.Option{
        validator.WithRequiredStructEnabled(),
    }

	walletRepositoryImpl := repository.NewWalletRepository()
	db := db.NewDB()
	validate := NewValidator(option)
	walletServiceImpl := service.NewWalletService(walletRepositoryImpl,db,validate)
	walletControllerImpl :=controller.NewWalletController(walletServiceImpl)

	var server = fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler() ,
	})
	SetupRoutes(server,walletControllerImpl)
	return server
}