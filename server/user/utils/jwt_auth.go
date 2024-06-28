package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Username  string  `json:"username"`
	Admin     bool    `json:"admin"`
	Key				string  `json:"key"`
	jwt.RegisteredClaims
}

func JWTAuth(usernameOrEmail string) string {
	if usernameOrEmail == "netsinx_15" || usernameOrEmail == "yasin@gmail.com" {
		signingKey := []byte("netsinxadmin")

		claims := CustomClaims{
			"netsinx_15",
			true,
			"jwtnetsinxadmin",
			jwt.RegisteredClaims{
				Issuer: "this is a jwt",
				IssuedAt: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			},
		}

		genToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		token, _ := genToken.SignedString(signingKey)

		return token
	}

	signingKey := []byte("yasinganteng15")

	claims := CustomClaims{
		usernameOrEmail,
		false,
		"jwtyasinganteng",
		jwt.RegisteredClaims{
			Issuer: "this is a jwt",
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	genToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, _ := genToken.SignedString(signingKey)

	return token
}