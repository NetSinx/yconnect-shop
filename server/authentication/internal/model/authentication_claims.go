package model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}
