package handlers

type JsonResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
	Body    string `json:"body"`
}
