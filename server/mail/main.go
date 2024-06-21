package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/NetSinx/yconnect-shop/server/mail/domain"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Please click this url: \n%v\n", url)

	var authToken string

	fmt.Print("Enter auth token: ")
	if _, err := fmt.Scan(&authToken); err != nil {
		log.Printf("Error message: %v", err)
	}

	token, err := config.Exchange(context.TODO(), authToken)
	if err != nil {
		return nil, err
	}

	openToken, err := os.OpenFile("token.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}
	defer openToken.Close()

	json.NewEncoder(openToken).Encode(token)

	return token, nil
}

func getTokenFromFile(nameFile string, token *oauth2.Token) (*oauth2.Token, error) {
	openToken, err := os.Open(nameFile)
	if err != nil {
		return nil, err
	}

	json.NewDecoder(openToken).Decode(token)

	return token, nil
}

func sendMail(ctx context.Context, config *oauth2.Config, token *oauth2.Token, receive string, otpCode string) error {
	client := config.Client(ctx, token)
	gmailService, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return err
	}

	messageStr := []byte(
		"From: yasin03ckm@gmail.com\r\n"+
		"To: "+ receive +"\r\n"+
		"Subject: Tester Mail\r\n"+
		"Content-Type: text/html; charset=utf-8\r\n\r\n"+
		"<h1>Welcome to Y-Connect Shop</h1>"+
		"<p>Your OTP Code: " + otpCode + "</p>")

	_, err = gmailService.Users.Messages.Send("me", &gmail.Message{Raw: base64.URLEncoding.EncodeToString(messageStr)}).Do()
	if err != nil {
		return err
	}

	return nil
}

func SendOTP(c echo.Context) error {
	var reqUser domain.ReqUser

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		json.NewEncoder(w).Encode(domain.ResponseMessage{
			Message: err.Error(),
		})
	}

	ctx := context.Background()

	secret, err := os.ReadFile("credentials.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ResponseMessage{
			Message: err.Error(),
		})

		return
	}

	conf, err := google.ConfigFromJSON(secret, gmail.GmailSendScope)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ResponseMessage{
			Message: err.Error(),
		})

		return
	}

	token := &oauth2.Token{}

	getToken, err := getTokenFromFile("token.json", token)
	if err != nil {
		token, err = getTokenFromWeb(conf)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(domain.ResponseMessage{
				Message: err.Error(),
			})

			return
		}

		if err := sendMail(ctx, conf, token, reqUser.Email, reqUser.OTP); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(domain.ResponseMessage{
				Message: err.Error(),
			})

			return
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(domain.ResponseMessage{
				Message: "Kode OTP berhasil dikirim.",
			})
		}

		return
	}

	if err := sendMail(ctx, conf, getToken, reqUser.Email, reqUser.OTP); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ResponseMessage{
			Message: err.Error(),
		})

		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(domain.ResponseMessage{
			Message: "Kode OTP berhasil dikirim.",
		})
	}

	return
}

func main() {
	router := echo.New()
	router.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:xsrf",
		CookieName: "xsrf",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
		CookieMaxAge: 60,
		CookieSecure: true,
	}))
	fmt.Println("Server running on localhost:8085...")
	log.Fatal(http.ListenAndServe(":8085", SendOTP()))
}