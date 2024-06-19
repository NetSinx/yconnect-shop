package test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"github.com/NetSinx/yconnect-shop/server/user/model/domain"
)

func TestListUser(t *testing.T) {
	response, err := http.Get("http://localhost:8082/user")

	if response.StatusCode != 200 {
		var respData domain.RespData

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, err.Error())
	}
}

func TestRegisterUser(t *testing.T) {
	body := `{
		"name": "Yasin Ayatulloh Hakim",
		"username": "netsinx_15",
		"email": "yasin@gmail.com",
		"alamat": "Jl. Kayu Manis",
		"no_telp": "087878504814",
		"password": "@Abyasinah22123"
	}`

	response, err := http.Post("http://localhost:8082/user/sign-up", "application/json", strings.NewReader(body))

	if response.StatusCode != 200 {
		var respData domain.RespData

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, err.Error())
	}
}

func TestUpdateUser(t *testing.T) {
	var httpClient http.Client

	body := `{
		"name": "Nur Azizah",
		"username": "azizah6",
		"email": "nurazizah@gmail.com",
		"alamat": "Jl. Kayu Manis",
		"no_telp": "089676798686",
		"password": "@Abnurazizah123"
	}`

	req, _ := http.NewRequest("PUT", "http://localhost:8082/user/5", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	response, err := httpClient.Do(req)

	if response.StatusCode != 200 {
		var respData domain.RespData

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, err.Error())
	}
}

func TestDeleteUser(t *testing.T) {
	var httpClient http.Client

	req, _ := http.NewRequest("DELETE", "http://localhost:8082/user/3", nil)

	response, err := httpClient.Do(req)

	if response.StatusCode != 200 {
		var respData domain.RespData

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, err.Error())
	}
}

func TestLoginUser(t *testing.T) {
	body := `{
		"email": "yasin@gmail.com",
		"password": "@Abyasinah22123"
	}`

	response, err := http.Post("http://localhost:8082/user/sign-in", "application/json", strings.NewReader(body))

	if response.StatusCode != 200 {
		var respData domain.RespData

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, err.Error())
	}
}

func TestGetUser(t *testing.T) {
	response, err := http.Get("http://localhost:8082/user/1")

	if response.StatusCode != 200 {
		var respData domain.RespData

		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %d, Error Message: %s", response.StatusCode, err.Error())
	}
}