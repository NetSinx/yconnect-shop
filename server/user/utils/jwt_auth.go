package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Username  string  `json:"username"`
	Role      string  `json:"role"`
	jwt.RegisteredClaims
}

func JWTAuth(username, role string) string {
	if role == "admin" {
		signingKey := []byte("netsinxadmin")

		claims := CustomClaims{
			username,
			role,
			jwt.RegisteredClaims{
				Issuer: "this is a jwt",
				IssuedAt: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			},
		}

		genToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		token, _ := genToken.SignedString(signingKey)

		return token
	} else {
		signingKey := []byte("yasinganteng15")
	
		claims := CustomClaims{
			username,
			role,
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
}