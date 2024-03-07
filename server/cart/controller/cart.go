package controller

import (
	"net/http"
	"strconv"
	"github.com/NetSinx/yconnect-shop/server/cart/model"
	"github.com/NetSinx/yconnect-shop/server/cart/service"
	"github.com/NetSinx/yconnect-shop/server/cart/utils"
	"github.com/labstack/echo/v4"
)

type cartController struct {
	cartService service.CartServ
}

func CartController(cs service.CartServ) cartController {
	return cartController{
		cartService: cs,
	}
}

func (cart cartController) ListCart(c echo.Context) error {
	var carts []model.Cart

	listCart, err := cart.cartService.ListCart(carts)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: listCart,
	})
}

func (cart cartController) AddToCart(c echo.Context) error {
	var cartModel model.Cart

	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&cartModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	}

	addCart, err := cart.cartService.AddToCart(cartModel, id)
	if err != nil && err.Error() == "request tidak sesuai" {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	} else if err != nil && err.Error() == "produk sudah ada" {
		return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
			Code: http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Message: "Produk sudah ada di keranjang!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: addCart,
	})
}

func (cart cartController) UpdateCart(c echo.Context) error {
	var cartModel model.Cart

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := c.Bind(&cartModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	}

	updCart, err := cart.cartService.UpdateCart(cartModel, uint(id))
	if err != nil && err.Error() == "request tidak sesuai" {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Produk tidak ditemukan!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: updCart,
	})
}

func (cart cartController) DeleteProductInCart(c echo.Context) error {
	var cartModel model.Cart

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := cart.cartService.DeleteProductInCart(cartModel, uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Produk tidak ditemukan di dalam keranjang",
		}) 
	}

	return c.JSON(http.StatusOK, utils.SuccessDelete{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Message: "Produk berhasil dihapus dari keranjang",
	})
}

func (cart cartController) GetCart(c echo.Context) error {
	var cartModel model.Cart

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	getCart, err := cart.cartService.GetCart(cartModel, uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Produk tidak ditemukan di dalam keranjang",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: getCart,
	})
}

func (cart cartController) GetCartByUser(c echo.Context) error {
	var cartModel []model.Cart

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	getCart, err := cart.cartService.GetCartByUser(cartModel, uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: getCart,
	})
}