package web

type SendRequest struct {
	FromAccount int    `validate:"required,minaccountid,numeric" json:"senderaccnumb"`
	ToAccount   int    `validate:"required,minaccountid,numeric" json:"recepientaccnumb"`
	Amount      float64 `validate:"required,numeric,min=10" json:"amount"`
}
