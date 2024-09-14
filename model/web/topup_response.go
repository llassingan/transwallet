package web

type TopUpResponse struct {
	TrxId     uint    `json:"idtrx"`
	AccountID uint    `json:"accnumb"`
	Amount    float64 `json:"amount"`
}
