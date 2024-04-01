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
		return echo.NewHTTPError(http.StatusInternalServerError, domain.ErrorResponse{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Sorry, there is available failure",
		})
	}

	return c.JSON(http.StatusOK, domain.FindAllResponse{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: listSeller,
	})
}