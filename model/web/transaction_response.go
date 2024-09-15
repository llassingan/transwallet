package web

import "time"

type TransactionResponse struct {
	TrxId           int      `json:"idtrx"`
	TransactionType string    `json:"trxtype"`
	Amount          float64   `json:"amount"`
	Time            time.Time `json:"time"`
}
