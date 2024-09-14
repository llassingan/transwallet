package web

type TopUpRequest struct {
	AccountID uint  `validate:"required,minaccountid" json:"accnumb"`
	Amount	float64 `validate:"required" json:"amount"`

}