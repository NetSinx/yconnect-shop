package model

type RegisterRequest struct {
	NamaLengkap        string `json:"nama_lengkap" validate:"required,max=100"`
	Username           string `json:"username" validate:"required,max=50"`
	Email              string `json:"email" validate:"required,max=100,email"`
	NoHP               string `json:"no_hp" validate:"required,max=12"`
	Password           string `json:"password" validate:"passwd,required,min=5"`
	KonfirmasiPassword string `json:"konfirmasi_password" validate:"required,eqfield=password"`
}

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

type AuthenticationResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type DataResponse struct {
	Data *AuthenticationResponse `json:"data"`
}
