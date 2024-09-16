package repository

import (
	"context"
	"transwallet/api/model/domain"
	"transwallet/api/model/web"

	"gorm.io/gorm"
)

type WalletRepository interface {
	// Top-up
	TopUp(ctx context.Context, tx *gorm.DB, accountId int, amount float64) (domain.Transaction, error)

	// Send Money
	SendMoney(ctx context.Context, tx *gorm.DB, fromAccountId int, toAccountId int, amount float64) (web.ReceiptResponse, error)

	// Get Balance
	GetBalance(ctx context.Context, tx *gorm.DB, accountId int) (domain.Account, error)

	// Get Transaction History
	GetTransactionHistory(ctx context.Context, tx *gorm.DB, accountId int) ([]domain.Transaction, error)
}