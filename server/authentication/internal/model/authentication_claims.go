package model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
