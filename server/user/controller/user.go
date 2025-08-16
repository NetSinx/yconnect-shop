package controller

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/NetSinx/yconnect-shop/server/user/model/domain"
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
	"github.com/NetSinx/yconnect-shop/server/user/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userController struct {
	userService service.UserServ
}

func UserController(userServ service.UserServ) userController {
	return userController{
		userService: userServ,
	}
}

func (u userController) RegisterUser(c echo.Context) error {
	var users entity.User

	if err := c.Bind(&users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	err := u.userService.RegisterUser(users)
	if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "User sudah terdaftar.",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Registrasi user berhasil.",
	})
}

func (u userController) ListUsers(c echo.Context) error {
	var users []entity.User

	listUsers, err := u.userService.ListUsers(users)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: listUsers,
	})
}

func (u userController) UpdateUser(c echo.Context) error {
	var user entity.User

	username := c.Param("username")
	email := c.Param("email")

	getDbUser, _ := u.userService.GetUser(user, username, email)
	
	avatar, err := c.FormFile("avatar")
	if err != nil {
		user.Avatar = ""

		os.Remove("." + getDbUser.Avatar)

		if err := c.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
				Message: err.Error(),
			})
		}

		err = u.userService.UpdateUser(user, username)
		if err != nil && err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
				Message: "User tidak ditemukan.",
			})
		} else if err != nil && err == gorm.ErrDuplicatedKey {
			return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
				Message: "User sudah pernah dibuat.",
			})
		} else if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, domain.MessageResp{
			Message: "User berhasil diupdate.",
		})
	}

	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileName := strings.Split(avatar.Filename, ".")[0]
	fileExt := strings.Split(avatar.Filename, ".")[1]
	hashedFileName := md5.New().Sum([]byte(fileName))

	os.MkdirAll("assets/images", os.ModePerm)

	dst, err := os.Create(fmt.Sprintf("assets/images/%x.%s", hashedFileName, fileExt))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	if getDbUser.Avatar != "" {
		os.Remove("." + getDbUser.Avatar)
	}

	user.Avatar = fmt.Sprintf("/assets/images/%x.%s", hashedFileName, fileExt)

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	err = u.userService.UpdateUser(user, username)
	if err != nil && err.Error() == "user tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil && err.Error() == "user sudah terdaftar" {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil && err == echo.ErrBadRequest {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "User berhasil diupdate.",
	})
}

func (u userController) VerifyOTP(c echo.Context) error {
	var verifyEmail domain.VerifyEmail

	if err := c.Bind(&verifyEmail); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	successMsg, err := u.userService.VerifyOTP(verifyEmail)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Email tidak sesuai dengan yang diverifikasi.",
		})
	} else if err != nil && err.Error() == "OTP tidak bisa dikirim" {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: successMsg,
	})
}

func (u userController) GetUser(c echo.Context) error {
	var users entity.User

	username := c.QueryParam("username")
	email := c.QueryParam("email")

	getUser, err := u.userService.GetUser(users, username, email)
	if err != nil && err.Error() == "user tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil && err == echo.ErrInternalServerError {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: getUser,
	})
}

func (u userController) VerifyEmail(c echo.Context) error {
	var verifyEmail domain.VerifyEmail

	if err := c.Bind(&verifyEmail); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	err := u.userService.VerifyEmail(verifyEmail)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil && err.Error() == "kode OTP tidak valid" {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return echo.NewHTTPError(http.StatusOK, domain.MessageResp{
		Message: "Email berhasil diverifikasi.",
	})
}

func (u userController) DeleteUser(c echo.Context) error {
	var users entity.User

	username := c.Param("username")
	email := c.Param("email")

	err := u.userService.DeleteUser(users, username, email)
	if err != nil && err.Error() == "user tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "User berhasil dihapus.",
	})
}
