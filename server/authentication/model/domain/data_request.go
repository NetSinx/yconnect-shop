package domain

type UserLogin struct {
	UsernameorEmail string `json:"UsernameorEmail" validate:"required"`
	Password        string `json:"password" validate:"required,min=5"`
}

type RequestTimezone struct {
	Timezone string `json:"timezone" validate:"required"`
}