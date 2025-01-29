package utils

import (
	"math/rand"
	"time"
)

func GenerateOTP() string {
	var token string
	strGenerator := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	
	rand.Seed(time.Now().Unix())

	for i := 0; i < 8; i++ {
		random := rand.Intn(len(strGenerator))
		token += string(strGenerator[random])
	}

	return token
}