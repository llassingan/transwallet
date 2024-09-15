package web

type BalanceResponse struct {
	AccountID int    `json:"accountnumber"`
	Balance    float64 `json:"balance"`
}
