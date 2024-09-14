package repository

import (
	"context"
	"errors"
	"transwallet/model/domain"
	"transwallet/model/web"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepositoryImpl struct {
}

// NewWalletRepository creates a new instance of WalletRepository
func NewWalletRepository(db *gorm.DB) *WalletRepositoryImpl {
	return &WalletRepositoryImpl{}
}

// TopUp adds funds to an account
func (r *WalletRepositoryImpl) TopUp(ctx context.Context, tx *gorm.DB, accountId int, amount float64) (domain.Transaction, error) {

	var account domain.Account
	var transaction domain.Transaction
	// perform locking to handle race condition
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&account, "id = ?", accountId).Error; err != nil {
		return transaction, err
	}

	// update balance
	account.Balance += amount
	if err := tx.Save(&account).Error; err != nil {
		return transaction, err
	}

	transaction = domain.Transaction{
		AccountID: uint(accountId),
		Amount:    amount,
		Type:      "c",
	}

	if err := tx.Create(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

// SendMoney transfers money between accounts
func (r *WalletRepositoryImpl) SendMoney(ctx context.Context, tx *gorm.DB, fromAccountId int, toAccountId int, amount float64) (web.ReceiptResponse, error) {
	var fromAccount, toAccount domain.Account
	var transactiondeb, transactioncred domain.Transaction

	// Fetch the accounts
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&fromAccount, "id = ?", fromAccountId).Error; err != nil {
		return web.ReceiptResponse{}, err
	}
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Preload("Customer").First(&toAccount, "id=?", toAccountId).Error; err != nil {
		return web.ReceiptResponse{}, err
	}

	// Check if the fromAccount has sufficient balance
	if fromAccount.Balance < amount {
		return web.ReceiptResponse{}, errors.New("insufficient funds")
	}

	// Update balances
	fromAccount.Balance -= amount
	toAccount.Balance += amount
	if err := tx.Save(&fromAccount).Error; err != nil {
		return web.ReceiptResponse{}, err
	}
	if err := tx.Save(&toAccount).Error; err != nil {
		return web.ReceiptResponse{}, err
	}

	// Create a new transaction record

	transactiondeb = domain.Transaction{
		AccountID: uint(toAccountId),
		Amount:    amount,
		Type:      "c",
	}
	if err := tx.Create(&transactiondeb).Error; err != nil {
		return web.ReceiptResponse{}, err
	}

	transactioncred = domain.Transaction{
		AccountID: uint(fromAccountId),
		Amount:    amount,
		Type:      "d",
	}
	if err := tx.Create(&transactioncred).Error; err != nil {
		return web.ReceiptResponse{}, err
	}
	// Prepare the response
	response := web.ReceiptResponse{
		IdTrx:            transactiondeb.ID,
		SenderAccNumb:    uint(fromAccountId),
		RecepientAccNumb: uint(toAccountId),
		RecepientName:    toAccount.Customer.Name,
		Amount:           amount,
	}

	return response, nil
}

// GetBalance returns the current balance of an account
func (r *WalletRepositoryImpl) GetBalance(ctx context.Context, tx *gorm.DB, accountId int) (domain.Account, error) {
	var account domain.Account

	// Fetch the account
	if err := tx.First(&account, "id= ?",accountId).Error; err != nil {
		return account, err
	}

	return account, nil
}

// GetTransactionHistory returns the transaction history for an account
func (r *WalletRepositoryImpl) GetTransactionHistory(ctx context.Context, tx *gorm.DB, accountId int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	// Fetch transactions for the account
	if err := tx.Where("account = ?", accountId).Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}
