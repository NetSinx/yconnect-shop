package model

type RegisterRequest struct {
	NamaLengkap        string `json:"nama_lengkap" validate:"required,max=100"`
	Username           string `json:"username" validate:"required,max=50"`
	Email              string `json:"email" validate:"required,max=100,email"`
	NoHP               string `json:"no_hp" validate:"required,max=12"`
	Password           string `json:"password" validate:"passwd,required,min=5"`
	KonfirmasiPassword string `json:"konfirmasi_password" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type AuthTokenRequest struct {
	AuthToken string `json:"auth_token" validate:"required"`
}

type LoginResponse struct {
	ID        uint   `json:"id"`
	Role      string `json:"role"`
	AuthToken string `json:"auth_token"`
}

type VerifyResponse struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
}

type AuthenticationResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type RegisterResponse struct {
	ID          uint   `json:"id"`
	NamaLengkap string `json:"nama_lengkap"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	NoHP        string `json:"no_hp"`
}

type DataResponse[T any] struct {
	Data T `json:"data"`
}
