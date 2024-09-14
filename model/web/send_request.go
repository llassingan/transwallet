package web

type SendRequest struct {
	FromAccount uint    `validate:"required,minaccountid" json:"senderaccnumb"`
	ToAccount   uint    `validate:"required,minaccountid" json:"recepientaccnumb"`
	Amount      float64 `validate:"required" json:"amount"`
}
