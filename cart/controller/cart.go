package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/NetSinx/yconnect-shop/cart/model"
	"github.com/NetSinx/yconnect-shop/cart/repository"
	"github.com/NetSinx/yconnect-shop/cart/utils"
	"github.com/labstack/echo/v4"
)

type cartController struct {
	cartRepository repository.CartRepo
}

func CartController(cartRepository repository.CartRepo) cartController {
	return cartController{
		cartRepository: cartRepository,
	}
}

func (cart cartController) ListCart(c echo.Context) error {
	var cartModel []model.Cart

	listCart, err := cart.cartRepository.ListCart(cartModel)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada beberapa kesalahan pada server",
		})
	}

	for i, cart := range listCart {
		var preloadUser utils.PreloadUser
		var preloadCategory utils.PreloadCategory
		var httpClient http.Client

		token := utils.GenerateToken("netsinx_15")

		reqCategory, _ := http.NewRequest("GET", fmt.Sprintf("http://kong-gateway:8000/categories/id/%d", cart.CategoryId), nil)

		reqUser, _ := http.NewRequest("GET", fmt.Sprintf("http://kong-gateway:8000/users/%d", cart.UserId), nil)
		reqUser.Header.Add("Authorization", token)

		resCategory, _ := httpClient.Do(reqCategory)
		
		resUser, err := httpClient.Do(reqUser)
		if err != nil {
			return c.JSON(http.StatusOK, utils.SuccessGet{
				Code: http.StatusOK,
				Status: http.StatusText(http.StatusOK),
				Data: listCart,
			})
		}

		if err := json.NewDecoder(resUser.Body).Decode(&preloadUser); err != nil {
			return err
		}

		if err := json.NewDecoder(resCategory.Body).Decode(&preloadCategory); err != nil {
			return err
		}

		listCart[i].Category = preloadCategory.Data
		listCart[i].User = preloadUser.Data
	}

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: listCart,
	})
}

func (cart cartController) AddToCart(c echo.Context) error {
	var cartModel model.Cart

	if err := c.Bind(&cartModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	}

	createCart, err := cart.cartRepository.AddToCart(cartModel)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
			Code: http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Message: "Produk sudah ada!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: createCart,
	})
}

func (cart cartController) UpdateCart(c echo.Context) error {
	var cartModel model.Cart

	getId := c.Param("id")

	if err := c.Bind(&cartModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	}

	updateCart, err := cart.cartRepository.UpdateCart(cartModel, getId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Produk tidak ditemukan di dalam keranjang",
		})
	}

	id, _ := strconv.ParseUint(getId, 10, 32)

	updateCart.Id = uint(id)

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: updateCart,
	})
}

func (cart cartController) DeleteProductInCart(c echo.Context) error {
	var cartModel model.Cart

	id := c.Param("id")

	if err := cart.cartRepository.DeleteProductInCart(cartModel, id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Produk tidak ditemukan di dalam keranjang",
		}) 
	}

	return c.JSON(http.StatusOK, utils.SuccessDelete{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Message: "Produk berhasil dihapus",
	})
}

func (cart cartController) GetCartById(c echo.Context) error {
	var cartModel model.Cart
	var httpClient http.Client
	var preloadCategory utils.PreloadCategory
	var preloadUser utils.PreloadUser

	id := c.Param("id")

	getCart, err := cart.cartRepository.GetCartById(cartModel, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Produk tidak ditemukan di dalam keranjang",
		})
	}

	token := utils.GenerateToken("netsinx_15")

	reqCategory, _ := http.NewRequest("GET", fmt.Sprintf("http://kong-gateway:8000/categories/id/%d", getCart.CategoryId), nil)

	reqUser, _ := http.NewRequest("GET", fmt.Sprintf("http://kong-gateway:8000/users/%d", getCart.UserId), nil)
	reqUser.Header.Add("Authorization", token)

	resCategory, _ := httpClient.Do(reqCategory)
	
	resUser, err := httpClient.Do(reqUser)
	if err != nil {
		return c.JSON(http.StatusOK, utils.SuccessGet{
			Code: http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data: getCart,
		})
	}

	json.NewDecoder(resCategory.Body).Decode(&preloadCategory)

	json.NewDecoder(resUser.Body).Decode(&preloadUser)

	getCart.Category = preloadCategory.Data
	getCart.User = preloadUser.Data

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: getCart,
	})
}

func (cart cartController) GetCartBySlug(c echo.Context) error {
	var cartModel model.Cart
	var httpClient http.Client
	var preloadCategory utils.PreloadCategory
	var preloadUser utils.PreloadUser

	slug := c.Param("slug")

	getCart, err := cart.cartRepository.GetCartBySlug(cartModel, slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Produk tidak ditemukan di dalam keranjang",
		}) 
	}

	token := utils.GenerateToken("netsinx_15")

	reqCategory, _ := http.NewRequest("GET", fmt.Sprintf("http://kong-gateway:8000/categories/id/%d", getCart.CategoryId), nil)

	reqUser, _ := http.NewRequest("GET", fmt.Sprintf("http://kong-gateway:8000/users/%d", getCart.UserId), nil)
	reqUser.Header.Add("Authorization", token)

	resCategory, _ := httpClient.Do(reqCategory)
	
	resUser, err := httpClient.Do(reqUser)
	if err != nil {
		return c.JSON(http.StatusOK, utils.SuccessGet{
			Code: http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data: getCart,
		})
	}

	json.NewDecoder(resCategory.Body).Decode(&preloadCategory)

	json.NewDecoder(resUser.Body).Decode(&preloadUser)

	getCart.Category = preloadCategory.Data
	getCart.User = preloadUser.Data

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: getCart,
	})
}