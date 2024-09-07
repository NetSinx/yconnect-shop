package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	prodEntity "github.com/NetSinx/yconnect-shop/server/product/model/entity"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
 reqOrder = entity.Order{
	Id: 1,
	ProductID: 1,
	Product: prodEntity.Product{
		Nama: "Baju Muslim",
		Slug: "baju-muslim",
		Gambar: append([]prodEntity.Gambar{}, prodEntity.Gambar{
			Id: 1,
			Nama: "baju_muslim.jpg",
			ProductID: 1,
		}),
		Deskripsi: "Baju nyaman dengan desain yang trendi",
		KategoriId: 1,
		Harga: 85000,
		Stok: 10,
		Rating: 4.8,
	},
	Username: "agus12",
	Kuantitas: 5,
	Status: "Dalam pengiriman",
	Estimasi: time.Now().AddDate(0, 0, 3),
 }

 successDelOrder = `{"message":"Pesanan berhasil dibatalkan"}`+"\n"
)

func TestListOrder(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/order", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:username")
	ctx.SetParamNames("username")
	ctx.SetParamValues("netsinx_15")
	respData, _ := json.Marshal(modelDB[ctx.Param("username")])

	if assert.NoError(t, ListOrder(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(respData)+"\n", rec.Body.String())
	}
}

func TestAddOrder(t *testing.T) {
	reqByte, _ := json.Marshal(reqOrder)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewReader(reqByte))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	reqOrder, _ := json.Marshal(reqOrder)

	if assert.NoError(t, AddOrder(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(reqOrder)+"\n", rec.Body.String())
	}
}

func TestDeleteOrder(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/order", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:username/:id")
	ctx.SetParamNames("username", "id")
	ctx.SetParamValues("netsinx_15", "1")

	if assert.NoError(t, DeleteOrder(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, successDelOrder, rec.Body.String())
	}
}