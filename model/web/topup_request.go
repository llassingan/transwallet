package web

type TopUpRequest struct {
	AccountID int  `validate:"required,minaccountid,numeric" json:"accnumb"`
	Amount	float64 `validate:"required,numeric,min=10" json:"amount"`

}