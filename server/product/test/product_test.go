package test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"github.com/NetSinx/yconnect-shop/server/product/model/domain"
)

func TestListProduct(t *testing.T) {
	response, _ := http.Get("http://localhost:8081/product")

	if response.StatusCode != 200 {
		var respData domain.MessageResp

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestCreateProduct(t *testing.T) {
	var genCSRF domain.ResponseCSRF
	var respData domain.MessageResp
	var httpClient http.Client
	var httpCookie http.Cookie

	body := `{
		"name": "Ayam Bakar",
		"slug": "ayam-bakar",
		"images": {
			"name": "ayam_bakar.jpg"
		},
		"description": "Ayam bakar sedap dengan cita rasa yang nikmat dan luar biasa",
		"category_id": 1,
		"seller_id": 1,
		"price": 19000,
		"stock": 10
	}`

	resp_gen_csrf, err := http.Get("http://localhost:8081/gencsrf")
	if resp_gen_csrf.StatusCode != 200 || err != nil {
		json.NewDecoder(resp_gen_csrf.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", resp_gen_csrf.StatusCode, respData.Message)
	}
	
	json.NewDecoder(resp_gen_csrf.Body).Decode(&genCSRF)
	
	request, _ := http.NewRequest("POST", "http://localhost:8081/product", strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("XSRF-Token", genCSRF.CSRFToken)

	httpCookie.Name = "_csrf"
	httpCookie.Value = genCSRF.CSRFToken
	request.AddCookie(&httpCookie)

	response, _ := httpClient.Do(request)

	if response.StatusCode != 200 {
		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s, csrf: %s, header: %s, cookie: %v", response.StatusCode, respData.Message,genCSRF.CSRFToken, request.Header.Get("XSRF-Token"), request.Cookies())
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
		var respData domain.MessageResp

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestDeleteProduct(t *testing.T) {
	var httpClient http.Client

	req, _ := http.NewRequest("DELETE", "http://localhost:8081/product/11", nil)

	response, _ := httpClient.Do(req)

	if response.StatusCode != 200 {
		var respData domain.MessageResp

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestGetProduct(t *testing.T) {
	response, _ := http.Get("http://localhost:8081/product/11")

	if response.StatusCode != 200 {
		var respData domain.MessageResp

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestGetCategoryProduct(t *testing.T) {
	response, _ := http.Get("http://localhost:8081/product/category/2")

	if response.StatusCode != 200 {
		var respData domain.MessageResp

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestGetUserProduct(t *testing.T) {
	response, _ := http.Get("http://localhost:8081/product/user/1")

	if response.StatusCode != 200 {
		var respData domain.MessageResp

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}