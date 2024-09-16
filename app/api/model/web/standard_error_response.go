package web

type StdErrorResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error   interface{} `json:"error"`
}