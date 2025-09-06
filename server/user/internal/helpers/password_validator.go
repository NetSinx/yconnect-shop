package helpers

import "github.com/go-playground/validator/v10"

func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	hasLetter := false
	hasNumber := false
	hasSymbol := false

	for _, c := range password {
		switch {
		case 'a' <= c && c <= 'z':
			hasLetter = true
		case 'A' <= c && c <= 'Z':
			hasLetter = true
		case '0' <= c && c <= '9':
			hasNumber = true
		default:
			hasSymbol = true
		}
	}

	return hasLetter && hasNumber && hasSymbol
}