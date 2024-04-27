package domain

type ReqUser struct {
	Email string `json:"email"`
}

type ResponseMessage struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}