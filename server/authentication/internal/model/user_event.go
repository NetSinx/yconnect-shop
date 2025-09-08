package model

type RegisterUserEvent struct {
	NamaLengkap string `json:"nama_lengkap"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	NoHP        string `json:"no_hp"`
	Role        string `json:"role"`
}