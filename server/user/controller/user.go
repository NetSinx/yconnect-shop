package controller

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/NetSinx/yconnect-shop/server/user/app/model"
	"github.com/NetSinx/yconnect-shop/server/user/service"
	"github.com/NetSinx/yconnect-shop/server/user/utils"
	"github.com/labstack/echo/v4"
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
	var users model.User

	if err := c.Bind(&users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	}

	err := u.userService.RegisterUser(users)
	if err != nil && err.Error() == "request tidak sesuai" {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	} else if err != nil && (err.Error() == "consumer gagal dibuat" || err.Error() == "token gagal dibuat") {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
			Code: http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Message: "User sudah pernah dibuat!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessCUD{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Message: "Registrasi user berhasil!",
	})
}

func (u userController) LoginUser(c echo.Context) error {
	var userLogin model.UserLogin

	if err := c.Bind(&userLogin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	}

	jwtToken, err := u.userService.LoginUser(userLogin)
	if err != nil && err.Error() == "email atau password salah" {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrServer{
			Code: http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Message: "Email atau password Anda salah!",
		})
	} else if err != nil && err.Error() == "email tidak mengandung karakter '@'" {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrServer{
			Code: http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Message: "Email atau password Anda salah!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{} {
		"code": http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"token": jwtToken,
	})
}

func (u userController) ListUsers(c echo.Context) error {
	var users []model.User
	
	listUsers, err := u.userService.ListUsers(users)
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
		Data: listUsers,
	})
}

func (u userController) UpdateUser(c echo.Context) error {
	var users model.User

	id := c.Param("id")

	getUser, _ := u.userService.GetUser(users, id)
	
	avatar, err := c.FormFile("avatar")
	if err != nil {
		users.Avatar = ""

		os.Remove("." + getUser.Avatar)

		err = u.userService.UpdateUser(users, id)
		if err != nil && err.Error() == "request tidak sesuai" {
			return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
				Code: http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Message: "Request yang dikirim tidak sesuai!",
			})
		} else if err != nil && err.Error() == "user tidak ditemukan" {
			return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
				Code: http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Message: "User tidak ditemukan!",
			})
		} else if err != nil && err.Error() == "user sudah pernah dibuat" {
			return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
				Code: http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Message: "User sudah pernah dibuat!",
			})
		}

		return c.JSON(http.StatusOK, utils.SuccessCUD{
			Code: http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Message: "User berhasil diupdate!",
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

	if getUser.Avatar != "" {
		os.Remove("." + getUser.Avatar)
	}

	users.Avatar = fmt.Sprintf("/assets/images/%x.%s", hashedFileName, fileExt)

	if err := c.Bind(&users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	}
	
	err = u.userService.UpdateUser(users, id)
	if err != nil && err.Error() == "user tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "User tidak ditemukan!",
		})
	} else if err != nil && err.Error() == "user sudah pernah dibuat" {
		return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
			Code: http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Message: "User sudah pernah dibuat!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	} 

	return c.JSON(http.StatusOK, utils.SuccessCUD{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Message: "User berhasil diupdate!",
	})
}

func (u userController) GetUser(c echo.Context) error {
	var users model.User

	id := c.Param("id")

	findUser, err := u.userService.GetUser(users, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "User tidak ditemukan!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: findUser,
	})
}

func (u userController) GetSeller(c echo.Context) error {
	var users model.User

	id := c.Param("id")

	getSeller, err := u.userService.GetSeller(users, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Seller tidak ditemukan!",
		}) 
	}

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: getSeller,
	})
}

func (u userController) DeleteUser(c echo.Context) error {
	var users model.User

	id := c.Param("id")

	user, _ := u.userService.GetUser(users, id)

	if user.Avatar != "" {
		os.Remove("." + user.Avatar)
	}

	err := u.userService.DeleteUser(users, id)
	if err != nil && err.Error() == "gagal hapus user" {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "User tidak ditemukan!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessCUD{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Message: "User berhasil dihapus!",
	})
}