package web

type TopUpResponse struct {
	TrxId     int    `json:"idtrx"`
	AccountID int    `json:"accnumb"`
	Amount    float64 `json:"amount"`
}
