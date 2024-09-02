package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	prodEntity "github.com/NetSinx/yconnect-shop/server/product/model/entity"
	"github.com/NetSinx/yconnect-shop/server/order/model/domain"
	"time"
)

func TestListOrder(t *testing.T) {
	resp, err := http.Get("http://localhost:8084/order/netsinx_15")
	if err != nil {
		t.Fatalf("Error message: %v", err)
	} else if resp.StatusCode != 200 {
		var respErr interface{}
		json.NewDecoder(resp.Body).Decode(&respErr)
		t.Fatalf("Error status: %v, error message: %v", resp.Status, respErr)
	}

	var respData domain.DataResp
	json.NewDecoder(resp.Body).Decode(&respData)
	t.Log(respData)
}

func TestAddOrder(t *testing.T) {
	gambar := []prodEntity.Gambar{}
	gambar = append(gambar, prodEntity.Gambar{Nama: "baju_muslim.jpg"})

	reqData := entity.Order{
		ProductID: 1,
		Product: prodEntity.Product{
			Nama: "Baju Muslim",
			Slug: "baju-muslim",
			Gambar: gambar,
			Deskripsi: "Baju untuk orang muslim",
			KategoriId: 1,
			Harga: 15000,
			Stok: 10,
			Rating: 4.8,
		},
		Username: "netsinx_15",
		Kuantitas: 5,
		Status: "Sedang Dikirim",
		Estimasi: time.Now().AddDate(0, 0 , 3),
	}

	data, err := json.Marshal(reqData)
	if err != nil {
		t.Fatalf("Was error when marshalling JSON data!")
	}

	resp, err := http.Post("http://localhost:8084/order", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Error message: %v", err)
	} else if resp.StatusCode != 200 {
		var respErr interface{}

		json.NewDecoder(resp.Body).Decode(&respErr)

		t.Fatalf("Service is not response or status not ok with status error: %v and response error: %v", resp.Status, respErr)
	}

	var respData domain.DataResp
	json.NewDecoder(resp.Body).Decode(&respData)
	t.Log(respData)
}