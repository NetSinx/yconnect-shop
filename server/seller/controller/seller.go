package controller

import (
	"net/http"

	"github.com/NetSinx/yconnect-shop/server/seller/model/domain"
	"github.com/NetSinx/yconnect-shop/server/seller/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: listSeller,
	})
}

func (sc sellerController) RegisterSeller(c echo.Context) error {
	var seller domain.Seller

	username := c.Param("username")

	if err := c.Bind(&seller); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	regSeller, err := sc.SellerService.RegisterSeller(username, seller)
	if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "Seller sudah terdaftar",
		})
	} else if err != nil && err.Error() == "seller gagal registrasi. user tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: regSeller,
	})
}

func (sc sellerController) UpdateSeller(c echo.Context) error {
	username := c.Param("username")

	updSeller, err := sc.SellerService.UpdateSeller(username)
	if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "Seller sudah terdaftar",
		})
	} else if err != nil && err.Error() == "seller tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil && err.Error() == "service user sedang bermasalah" {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: updSeller,
	})
}

func (sc sellerController) DeleteSeller(c echo.Context) error {
	username := c.Param("username")

	err := sc.SellerService.DeleteSeller(username)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Seller berhasil dihapus",
	})
}

func (sc sellerController) GetSeller(c echo.Context) error {
	username := c.Param("username")

	getSeller, err := sc.SellerService.GetSeller(username)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: getSeller,
	})
}