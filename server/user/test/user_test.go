package test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

func TestListUser(t *testing.T) {
	response, err := http.Get("http://localhost:8082/user")

	if err != nil {
		t.Fatalf("Error message: %v", err)
	} else if response.StatusCode != 200 {
		var respData interface{}
	
		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %v, Error Message: %v", response.StatusCode, respData)
	}

	var respData domain.RespData
	json.NewDecoder(response.Body).Decode(&respData)
	t.Log(respData)
}

func TestRegisterUser(t *testing.T) {
	body := `{
		"name": "Yasin Ayatulloh Hakim",
		"avatar": "",
		"username": "netsinx_15",
		"email": "yasin@gmail.com",
		"role": "admin",
		"alamat": {
			"alamat_rumah": "Jl. Kayu Manis no.15",
			"rt": 9,
			"rw": 8,
			"kelurahan": "Balekambang",
			"kecamatan": "Kramat Jati",
			"kota": "Jakarta Timur",
			"kode_pos": 13530
		},
		"no_telp": "087878504814",
		"password": "@Abyasinah22123",
		"email_verified": false
	}`

	response, err := http.Post("http://localhost:8082/user/sign-up", "application/json", strings.NewReader(body))

	if err != nil {
		t.Fatalf("Error message: %v", err)
	} else if response.StatusCode != 200 {
		var respData interface{}
		
		json.NewDecoder(response.Body).Decode(&respData)

		t.Fatalf("Error Status Code: %v, Error Message: %v", response.StatusCode, respData)
	}
}

func TestUpdateUser(t *testing.T) {
	var httpClient http.Client

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