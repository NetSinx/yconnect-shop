package controller

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/seller/model/domain"
	"github.com/NetSinx/yconnect-shop/server/seller/service"
	"github.com/labstack/echo/v4"
)

type sellerController struct {
	SellerService service.SellerServ
}

func SellerController(ss service.SellerServ) sellerController {
	return sellerController{
		SellerService: ss,
	}
}

func (sc sellerController) ListSeller(c echo.Context) error {
	listSeller, err := sc.SellerService.ListSeller()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.Response{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Sorry, there is available failure",
		})
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: listSeller,
	})
}

func (sc sellerController) RegisterSeller(c echo.Context) error {
	username := c.Param("username")

	regSeller, err := sc.SellerService.RegisterSeller(username)
	if err != nil && err.Error() == "seller sudah terdaftar" {
		return echo.NewHTTPError(http.StatusConflict, domain.Response{
			Code: http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Message: err.Error(),
		})
	} else if err != nil && err.Error() == "seller gagal registrasi. user tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.Response{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: regSeller,
	})
}

func (sc sellerController) UpdateSeller(c echo.Context) error {
	username := c.Param("username")

	updSeller, err := sc.SellerService.UpdateSeller(username)
	if err != nil && err.Error() == "seller sudah terdaftar" {
		return echo.NewHTTPError(http.StatusConflict, domain.Response{
			Code: http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Message: err.Error(),
		})
	} else if err != nil && err.Error() == "seller tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.Response{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: err.Error(),
		})
	} else if err != nil && err.Error() == "service user sedang bermasalah" {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.Response{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: updSeller,
	})
}

func (sc sellerController) DeleteSeller(c echo.Context) error {
	username := c.Param("username")

	if err := sc.SellerService.DeleteSeller(username); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.Response{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.Response{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Message: "Seller was deleted successfully",
	})
}

func (sc sellerController) GetSeller(c echo.Context) error {
	username := c.Param("username")

	getSeller, err := sc.SellerService.GetSeller(username)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.Response{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: getSeller,
	})
}