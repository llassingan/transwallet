package repository

import (
	"context"
	"errors"
	"transwallet/model/domain"
	"transwallet/model/web"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepositoryImpl struct {
	Logger *logrus.Logger
}

// NewWalletRepository creates a new instance of WalletRepository
func NewWalletRepository(logger *logrus.Logger) *WalletRepositoryImpl {
	return &WalletRepositoryImpl{
		Logger: logger,
	}
}

// TopUp adds funds to an account
func (r *WalletRepositoryImpl) TopUp(ctx context.Context, tx *gorm.DB, accountId uint, amount float64) (domain.Transaction, error) {

	var account domain.Account
	var transaction domain.Transaction
	// perform locking to handle race condition
	r.Logger.Info("Execute top up select acount query")
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&account, "id = ?", accountId).Error; err != nil {
		r.Logger.Error(err)
		return transaction,err
	}

	// update balance
	
	account.Balance += amount
	r.Logger.Info("Execute top up update account query")
	if err := tx.Save(&account).Error; err != nil {
		r.Logger.Error(err)
		return transaction, err
	}

	transaction = domain.Transaction{
		AccountID: accountId,
		Amount:    amount,
		Type:      "c",
	}
	r.Logger.Info("Execute top up insert debt transaction query")
	if err := tx.Create(&transaction).Error; err != nil {
		r.Logger.Error(err)
		return transaction, err
	}

	return transaction, nil
}

// SendMoney transfers money between accounts
func (r *WalletRepositoryImpl) SendMoney(ctx context.Context, tx *gorm.DB, fromAccountId uint, toAccountId uint, amount float64) (web.ReceiptResponse, error) {
	var fromAccount, toAccount domain.Account
	var transactiondeb, transactioncred domain.Transaction

	// Fetch the accounts
	r.Logger.Info("Execute send money select sender acount query")
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&fromAccount, "id = ?", fromAccountId).Error; err != nil {
		r.Logger.Error(err)
		return web.ReceiptResponse{}, err
	}
	r.Logger.Info("Execute send money select recepient acount query")
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Preload("Customer").First(&toAccount, "id=?", toAccountId).Error; err != nil {
		r.Logger.Error(err)
		return web.ReceiptResponse{}, err
	}

	// Check if the fromAccount has sufficient balance
	if fromAccount.Balance < amount {
		r.Logger.Error("insufficient funds")
		return web.ReceiptResponse{}, errors.New("insufficient funds")
	}

	// Update balances
	fromAccount.Balance -= amount
	toAccount.Balance += amount
	r.Logger.Info("Execute send money update sender acount query")
	if err := tx.Save(&fromAccount).Error; err != nil {
		r.Logger.Error(err)
		return web.ReceiptResponse{}, err
	}
	r.Logger.Info("Execute send money update recepient acount query")
	if err := tx.Save(&toAccount).Error; err != nil {
		r.Logger.Error(err)
		return web.ReceiptResponse{}, err
	}

	// Create a new transaction record
	
	transactiondeb = domain.Transaction{
		AccountID: toAccountId,
		Amount:    amount,
		Type:      "c",
	}
	r.Logger.Info("Execute send money insert credit transaction query")
	if err := tx.Create(&transactiondeb).Error; err != nil {
		r.Logger.Error(err)
		return web.ReceiptResponse{}, err
	}

	transactioncred = domain.Transaction{
		AccountID: fromAccountId,
		Amount:    amount,
		Type:      "d",
	}
	r.Logger.Info("Execute send money insert debt transaction query")
	if err := tx.Create(&transactioncred).Error; err != nil {
		r.Logger.Error(err)
		return web.ReceiptResponse{}, err
	}
	// Prepare the response
	response := web.ReceiptResponse{
		IdTrx:            transactiondeb.ID,
		SenderAccNumb:    fromAccountId,
		RecepientAccNumb: toAccountId,
		RecepientName:    toAccount.Customer.Name,
		Amount:           amount,
	}

	return response, nil
}

// GetBalance returns the current balance of an account
func (r *WalletRepositoryImpl) GetBalance(ctx context.Context, tx *gorm.DB, accountId uint) (domain.Account, error) {
	var account domain.Account

	// Fetch the account
	r.Logger.Info("Execute get balance select account query")
	if err := tx.First(&account, "id= ?",accountId).Error; err != nil {
		r.Logger.Error(err)
		return account, err
	}

	return account, nil
}

// GetTransactionHistory returns the transaction history for an account
func (r *WalletRepositoryImpl) GetTransactionHistory(ctx context.Context, tx *gorm.DB, accountId uint) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	// Fetch transactions for the account
	r.Logger.Info("Execute get history select query")
	if err := tx.Where("account = ?", accountId).Find(&transactions).Error; err != nil {
		r.Logger.Error(err)
		return transactions, err
	}

	return transactions, nil
}
