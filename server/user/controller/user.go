package controller

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
	"github.com/NetSinx/yconnect-shop/server/user/service"
	"github.com/NetSinx/yconnect-shop/server/user/model/domain"
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
	var users entity.User
	users.EmailVerified = false

	avatar, err := c.FormFile("avatar")
	if err != nil {
		if err := c.Bind(&users); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, domain.ErrServer{
				Code: http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Message: "Request yang dikirim tidak sesuai!",
			})
		}
	
		err := u.userService.RegisterUser(users)
		if err != nil && err.Error() == "request tidak sesuai" {
			return echo.NewHTTPError(http.StatusBadRequest, domain.ErrServer{
				Code: http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Message: "Request yang dikirim tidak sesuai!",
			})
		} else if err != nil && (err.Error() == "consumer gagal dibuat" || err.Error() == "token gagal dibuat") {
			return echo.NewHTTPError(http.StatusInternalServerError, domain.ErrServer{
				Code: http.StatusInternalServerError,
				Status: http.StatusText(http.StatusInternalServerError),
				Message: "Maaf, ada kesalahan pada server!",
			})
		} else if err != nil {
			return echo.NewHTTPError(http.StatusConflict, domain.ErrServer{
				Code: http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Message: "User sudah pernah dibuat!",
			})
		}
	
		return c.JSON(http.StatusOK, domain.SuccessCUD{
			Code: http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Message: "Registrasi user berhasil!",
		})
	}

	if err := os.MkdirAll("assets/images", os.ModePerm); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	src, _ := avatar.Open()
	dst, _ := os.Create(fmt.Sprintf("assets/images/%v", avatar.Filename))
	io.Copy(dst, src)

	if err := c.Bind(&users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	}

	err = u.userService.RegisterUser(users)
	if err != nil && err.Error() == "request tidak sesuai" {
		return echo.NewHTTPError(http.StatusBadRequest, domain.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	} else if err != nil && (err.Error() == "consumer gagal dibuat" || err.Error() == "token gagal dibuat") {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusConflict, domain.ErrServer{
			Code: http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Message: "User sudah pernah dibuat!",
		})
	}

	return c.JSON(http.StatusOK, domain.SuccessCUD{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Message: "Registrasi user berhasil!",
	})
}

func (u userController) LoginUser(c echo.Context) error {
	var userLogin entity.UserLogin

	if err := c.Bind(&userLogin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	}

	jwtToken, err := u.userService.LoginUser(userLogin)
	if err != nil && err.Error() == "email atau password salah" {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.ErrServer{
			Code: http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Message: "Email atau password Anda salah!",
		})
	} else if err != nil && err.Error() == "email tidak mengandung karakter '@'" {
		return echo.NewHTTPError(http.StatusBadRequest, domain.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.ErrServer{
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
	var users []entity.User
	
	listUsers, err := u.userService.ListUsers(users)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server",
		})
	}

	return c.JSON(http.StatusOK, domain.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: listUsers,
	})
}

func (u userController) UpdateUser(c echo.Context) error {
	var users entity.User

	username := c.Param("username")

	getUser, _ := u.userService.GetUser(users, username)
	
	avatar, err := c.FormFile("avatar")
	if err != nil {
		users.Avatar = ""

		os.Remove("." + getUser.Avatar)

		if err := c.Bind(&users); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, domain.ErrServer{
				Code: http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Message: "Request yang dikirim tidak sesuai!",
			})
		}

		err = u.userService.UpdateUser(users, username)
		if err != nil && err.Error() == "request tidak sesuai" {
			return echo.NewHTTPError(http.StatusBadRequest, domain.ErrServer{
				Code: http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Message: "Request yang dikirim tidak sesuai!",
			})
		} else if err != nil && err.Error() == "user tidak ditemukan" {
			return echo.NewHTTPError(http.StatusNotFound, domain.ErrServer{
				Code: http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Message: "User tidak ditemukan!",
			})
		} else if err != nil && err.Error() == "user sudah pernah dibuat" {
			return echo.NewHTTPError(http.StatusConflict, domain.ErrServer{
				Code: http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Message: "User sudah pernah dibuat!",
			})
		}

		return c.JSON(http.StatusOK, domain.SuccessCUD{
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
		return echo.NewHTTPError(http.StatusBadRequest, domain.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	}
	
	err = u.userService.UpdateUser(users, username)
	if err != nil && err.Error() == "user tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "User tidak ditemukan!",
		})
	} else if err != nil && err.Error() == "user sudah pernah dibuat" {
		return echo.NewHTTPError(http.StatusConflict, domain.ErrServer{
			Code: http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Message: "User sudah pernah dibuat!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	} 

	return c.JSON(http.StatusOK, domain.SuccessCUD{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Message: "User berhasil diupdate!",
	})
}

func (u userController) VerifyEmail(c echo.Context) error {
	var verifyEmail entity.VerifyEmail

	if err := c.Bind(&verifyEmail); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request tidak valid",
		})
	}

	token, err := u.userService.VerifyEmail(verifyEmail)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, domain.ErrServer{
			Code: http.StatusForbidden,
			Status: http.StatusText(http.StatusForbidden),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (u userController) GetUser(c echo.Context) error {
	var users entity.User

	username := c.Param("username")

	findUser, err := u.userService.GetUser(users, username)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "User tidak ditemukan!",
		})
	}

	return c.JSON(http.StatusOK, domain.SuccessGet{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: findUser,
	})
}

func (u userController) DeleteUser(c echo.Context) error {
	var users entity.User

	username := c.Param("username")

	user, _ := u.userService.GetUser(users, username)

	if user.Avatar != "" {
		os.Remove("." + user.Avatar)
	}

	err := u.userService.DeleteUser(users, username)
	if err != nil && err.Error() == "gagal hapus user" {
		return echo.NewHTTPError(http.StatusNotFound, domain.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "User tidak ditemukan!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server",
		})
	}

	return c.JSON(http.StatusOK, domain.SuccessCUD{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Message: "User berhasil dihapus!",
	})
}