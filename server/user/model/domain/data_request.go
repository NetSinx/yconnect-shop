package domain

import "database/sql"

type VerifyEmail struct {
	Email string `json:"email" validate:"required,email"`
	OTP		string `json:"otp" validate:"required"`
}

type EmailVerified struct {
	EmailVerified 	bool
	EmailVerifiedAt sql.NullTime
}

type RequestTimezone struct {
	Timezone string `json:"timezone" validate:"required"`
}