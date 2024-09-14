package web

type BalanceResponse struct {
	AccountID uint    `json:"accountnumber"`
	Balance    float64 `json:"balance"`
}
