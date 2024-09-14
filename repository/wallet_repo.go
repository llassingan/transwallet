package repository

import (
	"context"
	"transwallet/model/domain"
	"transwallet/model/web"

	"gorm.io/gorm"
)

type WalletRepository interface {
	// Top-up
	TopUp(ctx context.Context, tx *gorm.DB, accountId uint, amount float64) (domain.Transaction, error)

	// Send Money
	SendMoney(ctx context.Context, tx *gorm.DB, fromAccountId uint, toAccountId uint, amount float64) (web.ReceiptResponse, error)

	// Check Balance
	GetBalance(ctx context.Context, tx *gorm.DB, accountId uint) (domain.Account, error)

	// Check Transaction History
	GetTransactionHistory(ctx context.Context, tx *gorm.DB, accountId uint) ([]domain.Transaction, error)
}