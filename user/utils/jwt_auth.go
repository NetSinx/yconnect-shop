package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuth() string {
	signingKey := []byte("yasinganteng15")

	claims := &jwt.RegisteredClaims{
		Issuer: "this is a jwt",
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	}

	genToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, _ := genToken.SignedString(signingKey)

	return token
}