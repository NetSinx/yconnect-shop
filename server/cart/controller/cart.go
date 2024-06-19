package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/NetSinx/yconnect-shop/server/cart/model/domain"
	"github.com/NetSinx/yconnect-shop/server/cart/model/entity"
	"github.com/NetSinx/yconnect-shop/server/cart/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	var carts []entity.Cart

	listCart, err := cart.cartService.ListCart(carts)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: listCart,
	})
}

func (cart cartController) AddToCart(c echo.Context) error {
	var cartModel entity.Cart

	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&cartModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	addCart, err := cart.cartService.AddToCart(cartModel, id)
	if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "Produk sudah ada di keranjang.",
		})
	} else if err != nil && err.Error() == fmt.Sprintf("Error get product: %v", err) {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: addCart,
	})
}

func (cart cartController) UpdateCart(c echo.Context) error {
	var cartModel entity.Cart

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := c.Bind(&cartModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	updCart, err := cart.cartService.UpdateCart(cartModel, uint(id))
	if err != nil && err == gorm.ErrRecordNotFound{
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Produk tidak ditemukan.",
		})
	} else if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "Produk sudah ada.",
		})
	} else if err != nil && err.Error() == fmt.Sprintf("Error get product: %v", err) {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: updCart,
	})
}

func (cart cartController) DeleteProductInCart(c echo.Context) error {
	var cartModel entity.Cart

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	err := cart.cartService.DeleteProductInCart(cartModel, uint(id))
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Produk tidak ditemukan di dalam keranjang.",
		}) 
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Produk berhasil dihapus dari keranjang.",
	})
}

func (cart cartController) GetCart(c echo.Context) error {
	var cartModel entity.Cart

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	getCart, err := cart.cartService.GetCart(cartModel, uint(id))
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Produk tidak ditemukan di dalam keranjang.",
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: getCart,
	})
}

func (cart cartController) GetCartByUser(c echo.Context) error {
	var cartModel []entity.Cart

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	getCart, err := cart.cartService.GetCartByUser(cartModel, uint(id))
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Produk tidak ada di dalam keranjang.",
		})
	} else if err != nil && err.Error() == fmt.Sprintf("Error get product: %v", err){
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: getCart,
	})
}