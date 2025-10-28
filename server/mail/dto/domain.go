package domain

type ReqUser struct {
	Email string `json:"email"`
	OTP		string `json:"otp"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}