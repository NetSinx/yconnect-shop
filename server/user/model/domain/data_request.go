package domain

import "database/sql"

type VerifyEmail struct {
	Email string `json:"email" validate:"required,email"`
	OTP		string `json:"otp" validate:"required"`
}

type UserLogin struct {
	UsernameorEmail string `json:"UsernameorEmail" validate:"required"`
	Password        string `json:"password" validate:"required,min=5"`
}

type EmailVerified struct {
	EmailVerified 	bool
	EmailVerifiedAt sql.NullTime
}

type RequestTimezone struct {
	Timezone string `json:"timezone" validate:"required"`
}