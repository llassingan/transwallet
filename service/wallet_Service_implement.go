package service

import (
	"context"
	"transwallet/model/web"
	"transwallet/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WalletServiceImpl struct {
	DB       *gorm.DB
	Repo     repository.WalletRepository
	Validate *validator.Validate
	Logger   *logrus.Logger
}

// create a new instance of WalletService
func NewWalletService(repo repository.WalletRepository, db *gorm.DB, validate *validator.Validate, logger *logrus.Logger) *WalletServiceImpl {
	return &WalletServiceImpl{
		Repo:     repo,
		DB:       db,
		Validate: validate,
		Logger:   logger}
}

// handle top up request 
func (s *WalletServiceImpl) TopUp(ctx context.Context, req web.TopUpRequest) (web.TopUpResponse, error) {
	var TopUpResponse web.TopUpResponse
	
	// validate request
	err := s.Validate.Struct(req)
	if err != nil {
		s.Logger.Error(err)
		return TopUpResponse, err
	}
	// begin transaction
	tx := s.DB.Begin()
	s.Logger.Info("Begin transaction")

	// rollback
	defer tx.Rollback()

	// pass the request to repository layer
	res, err := s.Repo.TopUp(ctx, tx, req.AccountID, req.Amount)
	if err != nil {
		return TopUpResponse, err
	}
	// wrap the response 
	TopUpResponse = web.TopUpResponse{
		TrxId:     res.ID,
		AccountID: res.AccountID,
		Amount:    res.Amount,
	}
	// commit 
	tx.Commit()
	s.Logger.Info("Commit transaction")
	return TopUpResponse, nil

}

// handle send money request 
func (s *WalletServiceImpl) SendMoney(ctx context.Context, req web.SendRequest) (web.ReceiptResponse, error) {
	var ReceiptResponse web.ReceiptResponse
	err := s.Validate.Struct(req)
	if err != nil {
		s.Logger.Error(err)
		return ReceiptResponse, err
	}
	tx := s.DB.Begin()
	s.Logger.Info("Begin transaction")
	defer tx.Rollback()

	res, err := s.Repo.SendMoney(ctx, tx, req.FromAccount, req.ToAccount, req.Amount)
	if err != nil {
		return ReceiptResponse, err
	}

	
	ReceiptResponse = web.ReceiptResponse{
		IdTrx:            res.IdTrx,
		SenderAccNumb:    res.SenderAccNumb,
		RecepientAccNumb: res.RecepientAccNumb,
		RecepientName:    res.RecepientName,
		Amount:           res.Amount,
	}
	tx.Commit()
	s.Logger.Info("Commit transaction")
	return ReceiptResponse, nil
}

// handle get balance request 
func (s *WalletServiceImpl) GetBalance(ctx context.Context, accountId int) (web.BalanceResponse, error) {
	var BalanceResponse web.BalanceResponse
	tx := s.DB.Begin()
	s.Logger.Info("Begin transaction")
	defer tx.Rollback()

	res, err := s.Repo.GetBalance(ctx, tx, accountId)
	if err != nil {
		return BalanceResponse, err
	}
	
	BalanceResponse = web.BalanceResponse{
		AccountID: res.ID,
		Balance:   res.Balance,
	}
	tx.Commit()
	s.Logger.Info("Commit transaction")
	return BalanceResponse, nil
}

// handle get transaction history request 
func (s *WalletServiceImpl) GetTransactionHistory(ctx context.Context, accountId int) ([]web.TransactionResponse, error) {
	var ListHistory []web.TransactionResponse
	tx := s.DB.Begin()
	s.Logger.Info("Begin transaction")
	defer tx.Rollback()

	res, err := s.Repo.GetTransactionHistory(ctx, tx, accountId)
	if err != nil {
		return ListHistory, err
	}

	for _, transaction := range res {
		transactionResponse := web.TransactionResponse{
			TrxId:           transaction.ID,
			TransactionType: transaction.Type,
			Amount:          transaction.Amount,
			Time:            transaction.CreatedAt,
		}
		ListHistory = append(ListHistory, transactionResponse)
	}

	tx.Commit()
	s.Logger.Info("Commit transaction")
	return ListHistory, nil
}
