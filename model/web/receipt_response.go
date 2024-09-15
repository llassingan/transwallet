package web

type ReceiptResponse struct {
	IdTrx            int    `json:"idTrx"`
	SenderAccNumb    int    `json:"senderaccnumb"`
	RecepientAccNumb int    `json:"recepientaccnumb"`
	RecepientName    string  `json:"recepientname"`
	Amount           float64 `json:"amount"`
}
