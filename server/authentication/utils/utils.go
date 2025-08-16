package utils

import (
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Role      string  `json:"role"`
}

func GenerateAccessToken(username, email, role string) string {
		signingKey := []byte("yasinnetsinx15")

		claims := CustomClaims{
			jwt.RegisteredClaims{
				Issuer: "this is a jwt",
				IssuedAt: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			},
			username,
			email,
			role,
		}

		genToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		token, _ := genToken.SignedString(signingKey)

		return token
}

func GenerateRefreshToken(username, email, role string) string {
	signingKey := []byte("adminyasinnetsinx_15")

	claims := CustomClaims{
		jwt.RegisteredClaims{
			Issuer: "this is a jwt",
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
		username,
		email,
		role,
	}

	genToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, _ := genToken.SignedString(signingKey)

	return token
}

func SetCookies(c echo.Context, name string, value string, time time.Time) {
	cookie := http.Cookie{
		Name: name,
		Value: value,
		Expires: time,
		HttpOnly: true,
		Path: "/",
		SameSite: http.SameSiteStrictMode,
		Secure: true,
	}
	
	c.SetCookie(&cookie)
}