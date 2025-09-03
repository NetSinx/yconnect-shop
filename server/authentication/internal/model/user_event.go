package model

type UserEvent struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
