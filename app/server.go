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
	// set validator option
	option := []validator.Option{
        validator.WithRequiredStructEnabled(),
    }
	// init repo
	walletRepositoryImpl := repository.NewWalletRepository(log)
	// init db
	db := db.NewDB(log)
	// init validator
	validate := NewValidator(option)
	// init service
	walletServiceImpl := service.NewWalletService(walletRepositoryImpl,db,validate,log)
	// init controller
	walletControllerImpl :=controller.NewWalletController(walletServiceImpl,log)

	// create new fiber instance
	var server = fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler() ,
	})
	// setup routes
	SetupRoutes(server,walletControllerImpl)
	return server
}