package model

type LoginRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type AuthTokenRequest struct {
	AuthToken string `json:"auth_token" validate:"required"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type VerifyResponse struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
}

type AuthTokenResponse struct {
	AuthToken string `json:"auth_token"`
}
