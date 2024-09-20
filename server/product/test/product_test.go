package test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestListProduct(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/product", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	reqByte, _ := json.Marshal(productModel)

	if assert.NoError(t, ListProduct(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(reqByte)+"\n", rec.Body.String())
	}
}

func TestCreateProduct(t *testing.T) {
	expectedResp := `{"message":"Produk berhasil ditambahkan"}`+"\n"

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	writer.WriteField("nama", "Product Test")
	writer.WriteField("slug", "product-test")
	writer.WriteField("deskripsi", "ini hanyalah deskripsi product test")
	writer.WriteField("kategori_id", "2")
	writer.WriteField("harga", "50000")
	writer.WriteField("stok", "15")
	writer.WriteField("rating", "4.8")
	
	fileImages := []string{"laptop1.jpg", "laptop2.jpg", "laptop3.jpg"}

	for _, img := range fileImages {
		file, err := os.Open(img)
		assert.NoError(t, err)
		defer file.Close()
	
		part, err := writer.CreateFormFile("gambar", img)
		assert.NoError(t, err)
	
		_, err = io.Copy(part, file)
		assert.NoError(t, err)
	}
	
	writer.Close()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/product", body)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	
	if assert.NoError(t, CreateProduct(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedResp, rec.Body.String())
	}
}

func TestUpdateProduct(t *testing.T) {
	expectedResp := `{"message":"Produk berhasil diubah"}`+"\n"

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	writer.WriteField("nama", "Product Test 2")
	writer.WriteField("slug", "product-test-2")
	writer.WriteField("deskripsi", "ini hanyalah deskripsi product test")
	writer.WriteField("kategori_id", "2")
	writer.WriteField("harga", "50000")
	writer.WriteField("stok", "15")
	writer.WriteField("rating", "4.8")
	
	fileImages := []string{"laptop2.jpg"}

	for _, img := range fileImages {
		file, err := os.Open(img)
		assert.NoError(t, err)
		defer file.Close()
	
		part, err := writer.CreateFormFile("gambar", img)
		assert.NoError(t, err)
	
		_, err = io.Copy(part, file)
		assert.NoError(t, err)
	}
	
	writer.Close()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/product", body)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:slug")
	ctx.SetParamNames("slug")
	ctx.SetParamValues("product-test")
	
	if assert.NoError(t, UpdateProduct(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedResp, rec.Body.String())
	}
}

func TestDeleteProduct(t *testing.T) {
	expectedResp := `{"message":"Produk berhasil dihapus"}`+"\n"

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/product", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:slug")
	ctx.SetParamNames("slug")
	ctx.SetParamValues("product-test-2")

	if assert.NoError(t, DeleteProduct(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedResp, rec.Body.String())
	}
}

func TestGetProduct(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/product", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:slug")
	ctx.SetParamNames("slug")
	ctx.SetParamValues("product-test-2")

	respByte, _ := json.Marshal(productModel[2])

	if assert.NoError(t, GetProduct(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(respByte)+"\n", rec.Body.String())
	}
}

func TestGetCategoryProduct(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/product", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("2")

	respByte, _ := json.Marshal(productModel[2])

	if assert.NoError(t, GetProductByCategory(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "["+string(respByte)+"]\n", rec.Body.String())
	}
}
