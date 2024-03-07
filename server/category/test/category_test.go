package test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"github.com/NetSinx/yconnect-shop/server/category/utils"
)

func TestListCategory(t *testing.T) {
	response, _ := http.Get("http://localhost:8080/category")

	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)
		
		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestCreateCategory(t *testing.T) {
	body := `{
		"name": "Makanan",
		"slug": "makanan"
	}`

	response, _ := http.Post("http://localhost:8080/category", "application/json", strings.NewReader(body))
	if response.StatusCode != 200 {
		var respData utils.ErrServer

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestUpdateCategory(t *testing.T) {
	var httpClient http.Client

	body := `{
		"name": "Pakaian",
		"slug": "pakaian"
	}`

	req, _ := http.NewRequest("PUT", "http://localhost:8080/category/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	response, _ := httpClient.Do(req)

	if response.StatusCode != 200 {
		var respData utils.ErrServer
		
		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestDeleteCategory(t *testing.T) {
	var httpClient http.Client

	req, _ := http.NewRequest("DELETE", "http://localhost:8080/category/9", nil)

	response, _ := httpClient.Do(req)

	if response.StatusCode != 200 {
		var respData utils.ErrServer
		
		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}

func TestGetCategory(t *testing.T) {
	response, _ := http.Get("http://localhost:8080/category/1")

	if response.StatusCode != 200 {
		var respData utils.ErrServer
		
		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, respData.Message)
	}
}