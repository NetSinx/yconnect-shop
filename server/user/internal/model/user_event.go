package model

type RegisterUserEvent struct {
	NamaLengkap string    `json:"nama_lengkap"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	NoHP        string    `json:"no_hp"`
}

type DeleteUserEvent struct {
	Email string `json:"email"`
}
