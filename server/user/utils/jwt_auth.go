package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Username  string  `json:"username"`
	Role      string  `json:"role"`
}

func JWTAuth(username, role string) string {
		signingKey := []byte("yasinnetsinx15")

		claims := CustomClaims{
			jwt.RegisteredClaims{
				Issuer: "this is a jwt",
				IssuedAt: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			},
			username,
			role,
		}

		genToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		token, _ := genToken.SignedString(signingKey)

		return token
}