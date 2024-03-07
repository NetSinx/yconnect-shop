package test

import (
	"encoding/json"
	"net/http"
	"testing"
	"strings"
	"github.com/NetSinx/yconnect-shop/server/cart/utils"
)

func TestListCart(t *testing.T) {
	response, _ := http.Get("http://localhost:8083/cart")

	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestAddToCart(t *testing.T) {
	body := `{
		"name": "Ayam Goreng",
		"slug": "ayam-goreng",
		"price": 17000,
		"item": 5,
		"user_id": 1,
		"category_id": 1
	}`

	response, _ := http.Post("http://localhost:8083/cart/3", "application/json", strings.NewReader(body))

	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestDeleteProductInCart(t *testing.T) {
	var httpClient http.Client

	req, _ := http.NewRequest("DELETE", "http://localhost:8083/cart/1", nil)

	response, _ := httpClient.Do(req)

	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestGetCart(t *testing.T) {
	response, _ := http.Get("http://localhost:8083/cart/1")

	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}