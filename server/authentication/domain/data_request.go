package domain

type UserLogin struct {
	UsernameorEmail string `json:"UsernameorEmail" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
}