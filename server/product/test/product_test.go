package test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"github.com/NetSinx/yconnect-shop/server/product/utils"
)

func TestListProduct(t *testing.T) {
	response, _ := http.Get("http://localhost:8081/product")

	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestCreateProduct(t *testing.T) {
	body := `{
		"name": "Ayam Bakar",
		"slug": "ayam-bakar",
		"description": "Ayam bakar sedap dengan cita rasa yang nikmat dan luar biasa",
		"category_id": 1,
		"seller_id": 1,
		"price": 19000,
		"stock": 10
	}`

	response, _ := http.Post("http://localhost:8081/product", "application/json", strings.NewReader(body))

	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestUpdateProduct(t *testing.T) {
	var httpClient http.Client

	body := `{
		"name": "Ayam Bakar",
		"slug": "ayam-bakar",
		"description": "Ayam bakar sedap dengan cita rasa yang nikmat dan luar biasa",
		"category_id": 1,
		"seller_id": 1,
		"price": 17000,
		"stock": 10
	}`

	req, _ := http.NewRequest("PUT", "http://localhost:8081/product/11", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	response, _ := httpClient.Do(req)

	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestDeleteProduct(t *testing.T) {
	var httpClient http.Client

	req, _ := http.NewRequest("DELETE", "http://localhost:8081/product/11", nil)

	response, _ := httpClient.Do(req)

	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestGetProduct(t *testing.T) {
	response, _ := http.Get("http://localhost:8081/product/11")

	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestGetCategoryProduct(t *testing.T) {
	response, _ := http.Get("http://localhost:8081/product/category/2")

	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestGetUserProduct(t *testing.T) {
	response, _ := http.Get("http://localhost:8081/product/user/1")

	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}