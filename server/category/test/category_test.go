package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/NetSinx/yconnect-shop/server/category/model/entity"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestListCategory(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/category", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	cateModel, _ := json.Marshal(categoryModel)
	
	if assert.NoError(t, ListCategory(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(cateModel)+"\n", rec.Body.String())
	}
}

func TestCreateCategory(t *testing.T) {
	expectedResp := `{"message":"Kategori berhasil ditambahkan"}`+"\n"

	reqCategory := entity.Kategori{
		Name: "Hats",
		Slug: "hats",
	}

	byteReq, _ := json.Marshal(reqCategory)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/category", bytes.NewReader(byteReq))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, CreateCategory(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedResp, rec.Body.String())
	}
}

func TestUpdateCategory(t *testing.T) {
	reqCategory := entity.Kategori{
		Name: "Shoes",
		Slug: "shoes",
	}

	expectedResp := `{"message":"Kategori berhasil diubah"}`+"\n"

	byteReq, _ := json.Marshal(reqCategory)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/category", bytes.NewReader(byteReq))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("2")

	if assert.NoError(t, UpdateCategory(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedResp, rec.Body.String())
	}
}

func TestDeleteCategory(t *testing.T) {
	expectedResp := `{"message":"Kategori berhasil dihapus"}`+"\n"

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/category", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, DeleteCategory(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedResp, rec.Body.String())
	}
}

func TestGetCategory(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/category", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	respData, _ := json.Marshal(categoryModel["data"][0])

	if assert.NoError(t, GetCategory(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(respData)+"\n", rec.Body.String())
	}
}