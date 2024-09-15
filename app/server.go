package app

import (
	"transwallet/controller"
	"transwallet/db"
	"transwallet/repository"
	"transwallet/service"
	"transwallet/utils/exception"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewServer(log *logrus.Logger) *fiber.App {
	option := []validator.Option{
        validator.WithRequiredStructEnabled(),
    }
	
	walletRepositoryImpl := repository.NewWalletRepository(log)
	db := db.NewDB(log)
	validate := NewValidator(option)
	walletServiceImpl := service.NewWalletService(walletRepositoryImpl,db,validate,log)
	walletControllerImpl :=controller.NewWalletController(walletServiceImpl,log)

	var server = fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler() ,
	})
	SetupRoutes(server,walletControllerImpl)
	return server
}