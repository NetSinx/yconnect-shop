package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"time"
)

type orderHandler struct {
	db map[string]*entity.Order
}

var (
	reqOrder = `{"product_id": 1, "username": "netsinx_15", "kuantitas": 5, "status": "Dalam pengiriman", "estimasi": `+ time.Now() +`}`
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

	strings.NewReader()
	
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