package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	modelDB = map[string]*entity.Order{
	 "netsinx_15": &entity.Order{
		 ProductID: 1,
		 Username: "netsinx_15",
		 Kuantitas: 5,
		 Status: "Dalam pengiriman",
	 },
 }

 orderJSON = `{"product_id": 1, "username": "netsinx_15", "kuantitas": 5, "status": "Dalam pengiriman"}`

 successAddOrder = `{"message": "Pesanan berhasil dibuat"}`
)

func TestListOrder(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/order", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:username")
	ctx.SetParamNames("username")
	ctx.SetParamValues("netsinx_15")
	h := &orderHandler{db: modelDB}
	
	if assert.NoError(t, h.ListOrder(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, orderJSON, rec.Body.String())
	}
}

func TestAddOrder(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/order", strings.NewReader(orderJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	h := &orderHandler{db: modelDB}

	if assert.NoError(t, h.AddOrder(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, successAddOrder, rec.Body.String())
	}
}