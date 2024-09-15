package repository

import (
	"context"
	"errors"
	"transwallet/api/model/domain"
	"transwallet/api/model/web"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepositoryImpl struct {
	Logger *logrus.Logger
}

//  create a new instance of WalletRepository
func NewWalletRepository(logger *logrus.Logger) *WalletRepositoryImpl {
	return &WalletRepositoryImpl{
		Logger: logger,
	}
}

// handle top up request
func (r *WalletRepositoryImpl) TopUp(ctx context.Context, tx *gorm.DB, accountId int, amount float64) (domain.Transaction, error) {

	var account domain.Account
	var transaction domain.Transaction
	
	r.Logger.Info("Execute top up select acount query")
	// perform locking to handle race condition
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

	// create transaction value
	transaction = domain.Transaction{
		AccountID: accountId,
		Amount:    amount,
		Type:      "c",
	}
	r.Logger.Info("Execute top up insert debt transaction query")
	// insert transaction data
	if err := tx.Create(&transaction).Error; err != nil {
		r.Logger.Error(err)
		return transaction, err
	}

	return transaction, nil
}

// handle send money request
func (r *WalletRepositoryImpl) SendMoney(ctx context.Context, tx *gorm.DB, fromAccountId int, toAccountId int, amount float64) (web.ReceiptResponse, error) {
	var fromAccount, toAccount domain.Account
	var transactiondeb, transactioncred domain.Transaction

	// fetch the sender and recepient accounts
	r.Logger.Info("Execute send money select sender acount query")
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&fromAccount, "id = ?", fromAccountId).Error; err != nil {
		r.Logger.Error(err)
		return web.ReceiptResponse{}, errors.New("sender not found")
	}
	r.Logger.Info("Execute send money select recepient acount query")
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Preload("Customer").First(&toAccount, "id=?", toAccountId).Error; err != nil {
		r.Logger.Error(err)
		return web.ReceiptResponse{}, errors.New("recepient not found")
	}

	// check if the sender has sufficient balance
	if fromAccount.Balance < amount {
		r.Logger.Error("insufficient funds")
		return web.ReceiptResponse{}, errors.New("insufficient funds")
	}

	// update balances for both account
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

	// create a new transaction record
	
	transactiondeb = domain.Transaction{
		AccountID: toAccountId,
		Amount:    amount,
		Type:      "c",
	}
	r.Logger.Info("Execute send money insert credit transaction query")
	// insert transaction data for both user
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
	// wrap the response
	response := web.ReceiptResponse{
		IdTrx:            transactiondeb.ID,
		SenderAccNumb:    fromAccountId,
		RecepientAccNumb: toAccountId,
		RecepientName:    toAccount.Customer.Name,
		Amount:           amount,
	}

	return response, nil
}

// handle get balance request
func (r *WalletRepositoryImpl) GetBalance(ctx context.Context, tx *gorm.DB, accountId int) (domain.Account, error) {
	var account domain.Account

	// fetch the account
	r.Logger.Info("Execute get balance select account query")
	// get the balance data
	if err := tx.First(&account, "id= ?",accountId).Error; err != nil {
		r.Logger.Error(err)
		return account, err
	}

	return account, nil
}

// handle get transaction history
func (r *WalletRepositoryImpl) GetTransactionHistory(ctx context.Context, tx *gorm.DB, accountId int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	// fetch transactions data

	r.Logger.Info("Execute get history select query")
	if err := tx.Where("account = ?", accountId).Find(&transactions).Error; err != nil {
		r.Logger.Error(err)
		return transactions, err
	}

	return transactions, nil
}
