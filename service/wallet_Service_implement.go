package service

import (
	"context"
	"transwallet/model/web"
	"transwallet/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type WalletServiceImpl struct {
	DB       *gorm.DB
	Repo     repository.WalletRepository
	Validate *validator.Validate
}

func NewWalletService(repo repository.WalletRepository, db *gorm.DB, validate *validator.Validate) *WalletServiceImpl {
	return &WalletServiceImpl{
		Repo:     repo,
		DB:       db,
		Validate: validate}
}

func (s *WalletServiceImpl) TopUp(ctx context.Context, req web.TopUpRequest) (web.TopUpResponse, error) {
	var TopUpResponse web.TopUpResponse
	err := s.Validate.Struct(req)
	if err != nil {
		return TopUpResponse, err
	}
	tx := s.DB.Begin()
	defer tx.Rollback()

	res, err := s.Repo.TopUp(ctx, tx, req.AccountID, req.Amount)
	if err != nil {
		return TopUpResponse, err
	}
	tx.Commit()
	TopUpResponse = web.TopUpResponse{
		TrxId:     res.ID,
		AccountID: res.AccountID,
		Amount:    res.Amount,
	}
	return TopUpResponse, nil

}

func (s *WalletServiceImpl) SendMoney(ctx context.Context, req web.SendRequest) (web.ReceiptResponse, error) {
	var ReceiptResponse web.ReceiptResponse
	err := s.Validate.Struct(req)
	if err != nil {
		return ReceiptResponse, err
	}
	tx := s.DB.Begin()
	defer tx.Rollback()

	res, err := s.Repo.SendMoney(ctx, tx, req.FromAccount, req.ToAccount, req.Amount)
	if err != nil {
		return ReceiptResponse, err
	}

	tx.Commit()
	ReceiptResponse = web.ReceiptResponse{
		IdTrx:            res.IdTrx,
		SenderAccNumb:    res.SenderAccNumb,
		RecepientAccNumb: res.RecepientAccNumb,
		RecepientName:    res.RecepientName,
		Amount:           res.Amount,
	}

	return ReceiptResponse, nil
}

func (s *WalletServiceImpl) GetBalance(ctx context.Context, accountId uint) (web.BalanceResponse,error) {
	var BalanceResponse web.BalanceResponse
	tx := s.DB.Begin()
	defer tx.Rollback()

	res,err := s.Repo.GetBalance(ctx,tx,accountId)
	if err != nil {
		return BalanceResponse, err
	}
	tx.Commit()
	BalanceResponse = web.BalanceResponse{
		AccountID: res.ID,
		Balance: res.Balance,
	}

	return BalanceResponse, nil
}

func (s *WalletServiceImpl) GetTransactionHistory(ctx context.Context, accountId uint) ([]web.TransactionResponse,error) {
	var ListHistory []web.TransactionResponse
	tx := s.DB.Begin()
	defer tx.Rollback()

	res,err := s.Repo.GetTransactionHistory(ctx,tx,accountId)
	if err != nil {
		return ListHistory, err
	}

	for _, transaction := range res {
		transactionResponse := web.TransactionResponse{
			TrxId: transaction.ID,
			TransactionType: transaction.Type,
			Amount: transaction.Amount,
			Time: transaction.CreatedAt,
		}
		ListHistory = append(ListHistory, transactionResponse)
	}

	tx.Commit()

	return ListHistory,nil
}
