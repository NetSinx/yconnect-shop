package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"github.com/NetSinx/yconnect-shop/server/seller/model/domain"
)

func TestFindAllSeller(t *testing.T) {
	var response domain.RespData

	respFind, err := http.Get("http://localhost:8084/seller")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	json.NewDecoder(respFind.Body).Decode(&response)

	t.Logf("Find all seller successfully with data:\n%v", response)
}

func TestRegisterSeller(t *testing.T) {
	var responseSuccess domain.RespData
	var responseError domain.MessageResp

	reqRegistSeller := `{
		"name": "Yasin Ayatulloh Hakim",
		"username": "netsinx_15",
		"avatar": "",
		"email": "yasin123@gmail.com",
		"alamat": "Jl.Kayu Manis",
		"no_telp": "089676798686",
		"product": [],
		"user_id": 0
	}`

	resp, err := http.Post("http://localhost:8084/seller", "application/json", bytes.NewReader([]byte(reqRegistSeller)))
	if err != nil {
		t.Fatalf("Error: %v", err)
	} else if resp.StatusCode != 200 {
		json.NewDecoder(resp.Body).Decode(&responseError)
		t.Fatalf("Error: %v", responseError.Message)
	}
	
	json.NewDecoder(resp.Body).Decode(&responseSuccess)
	
	t.Logf("Post data successfully with data: %v", responseSuccess)
}