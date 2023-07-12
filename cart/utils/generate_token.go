package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) string {
	if username == "netsinx_15" {
		key := []byte("netsinxadmin")

		claims := CustomClaims{
			"netsinx_15",
			true,
			jwt.RegisteredClaims{
				Issuer: "this is a jwt",
				IssuedAt: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			},
		}

		newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		token, _ := newToken.SignedString(key)

		return token
	}

	key := []byte("yasinganteng15")

	claims := &jwt.RegisteredClaims{
		Issuer: "this is a jwt",
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, _ := newToken.SignedString(key)

	return token
}