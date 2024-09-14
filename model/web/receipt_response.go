package web

type ReceiptResponse struct {
	IdTrx            uint    `json:"idTrx"`
	SenderAccNumb    uint    `json:"senderaccnumb"`
	RecepientAccNumb uint    `json:"recepientaccnumb"`
	RecepientName    string  `json:"recepientname"`
	Amount           float64 `json:"amount"`
}
