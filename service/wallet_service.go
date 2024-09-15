package service

import (
	"context"
	"transwallet/model/web"
)

type WalletService interface {

	// Top-up
	TopUp(ctx context.Context, req web.TopUpRequest) (web.TopUpResponse, error)

	// Send Money
	SendMoney(ctx context.Context, req web.SendRequest) (web.ReceiptResponse, error)

	// Get Balance
	GetBalance(ctx context.Context, accountId int) (web.BalanceResponse, error)

	// Get Transaction History
	GetTransactionHistory(ctx context.Context, accountId int) ([]web.TransactionResponse, error)
}
