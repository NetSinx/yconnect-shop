package utils

import (
	"net/http"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var (
	AdminJwtKey = os.Getenv("JWT_KEY_ADMIN")
	CustomerJwtKey = os.Getenv("JWT_KEY_CUSTOMER")
)

type CustomClaims struct {
	Username  string  `json:"username"`
	Role      string  `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(username, role string) string {
	if role == "admin" {
		signingKey := []byte(AdminJwtKey)
		claims := CustomClaims{
			username,
			role,
			jwt.RegisteredClaims{
				Issuer: "this is a jwt",
				IssuedAt: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			},
		}

		genToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		token, _ := genToken.SignedString(signingKey)

		return token
	}

	signingKey := []byte(CustomerJwtKey)
	claims := CustomClaims{
		username,
		role,
		jwt.RegisteredClaims{
			Issuer: "this is a jwt",
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}

	genToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, _ := genToken.SignedString(signingKey)

	return token
}

func GenerateRefreshToken(username, role string) string {
	if role == "admin" {
		signingKey := []byte(AdminJwtKey)	
		claims := CustomClaims{
			username,
			role,
			jwt.RegisteredClaims{
				Issuer: "this is a jwt",
				IssuedAt: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			},
		}
	
		genToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		token, _ := genToken.SignedString(signingKey)
	
		return token
	}

	signingKey := []byte(CustomerJwtKey)
	claims := CustomClaims{
		username,
		role,
		jwt.RegisteredClaims{
			Issuer: "this is a jwt",
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
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