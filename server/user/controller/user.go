package controller

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/NetSinx/yconnect-shop/server/user/model/domain"
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
	"github.com/NetSinx/yconnect-shop/server/user/service"
	"github.com/golang-jwt/jwt/v4"
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
	users.EmailVerified = false

	avatar, err := c.FormFile("avatar")
	if err != nil {
		if err := c.Bind(&users); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
				Message: err.Error(),
			})
		}
	
		err := u.userService.RegisterUser(users)
		if err != nil && err == gorm.ErrDuplicatedKey {
			return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
				Message: "User sudah terdaftar",
			})
		} else if err != nil && (err.Error() == "consumer gagal dibuat" || err.Error() == "token gagal dibuat") {
			return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
				Message: err.Error(),
			})
		} else if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
				Message: err.Error(),
			})
		}
	
		return c.JSON(http.StatusOK, domain.MessageResp{
			Message: "Registrasi user berhasil!",
		})
	}

	if err := os.MkdirAll("assets/images", os.ModePerm); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	src, _ := avatar.Open()
	dst, _ := os.Create(fmt.Sprintf("assets/images/%v", avatar.Filename))
	io.Copy(dst, src)

	if err := c.Bind(&users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	err = u.userService.RegisterUser(users)
	if err != nil && (err.Error() == "consumer gagal dibuat" || err.Error() == "token gagal dibuat") {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "User sudah pernah dibuat!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Registrasi user berhasil!",
	})
}

func (u userController) LoginUser(c echo.Context) error {
	var userLogin entity.UserLogin

	if err := c.Bind(&userLogin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	jwtToken, err := u.userService.LoginUser(userLogin)
	if err != nil && err.Error() == "email atau password salah" {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: "Email atau password Anda salah!",
		})
	} else if err != nil && err.Error() == "email tidak mengandung karakter '@' dan hostname" {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: err.Error(),
		})
	}

	var cookie http.Cookie
	cookie.Name = "jwt_token"
	cookie.Value = jwtToken
	cookie.Expires = time.Now().Add(30 * time.Minute)
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Secure = true
	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, map[string]interface{} {
		"token": jwtToken,
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
	var users entity.User

	username := c.Param("username")

	getUser, _ := u.userService.GetUser(users, username)
	
	avatar, err := c.FormFile("avatar")
	if err != nil {
		users.Avatar = ""

		os.Remove("." + getUser.Avatar)

		if err := c.Bind(&users); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
				Message: err.Error(),
			})
		}

		err = u.userService.UpdateUser(users, username)
		if err != nil && err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
				Message: "User tidak ditemukan!",
			})
		} else if err != nil && err == gorm.ErrDuplicatedKey {
			return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
				Message: "User sudah pernah dibuat!",
			})
		} else if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, domain.MessageResp{
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
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}
	
	err = u.userService.UpdateUser(users, username)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "User tidak ditemukan!",
		})
	} else if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "User sudah pernah dibuat!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	} 

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "User berhasil diupdate!",
	})
}

func (u userController) SendOTP(c echo.Context) error {
	var verifyEmail domain.VerifyEmail

	if err := c.Bind(&verifyEmail); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	successMsg, err := u.userService.SendOTP(verifyEmail)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Email tidak sesuai dengan yang diverifikasi.",
		})
	}	else if err != nil && err.Error() == "OTP tidak bisa dikirim" {
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

	username := c.Param("username")

	findUser, err := u.userService.GetUser(users, username)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "User tidak ditemukan!",
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: findUser,
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
	if (err != nil && err == gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	} else if (err != nil && err.Error() == "kode OTP tidak valid") {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return echo.NewHTTPError(http.StatusOK, domain.MessageResp{
		Message: "Email berhasil diverifikasi!",
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
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "User tidak ditemukan!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "User berhasil dihapus!",
	})
}

func (u userController) IsLogin(c echo.Context) error {
	cookie, err := c.Cookie("jwt_token")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.RespData{
			Data: false,
		})
	} else if cookie.Value == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.RespData{
			Data: false,
		})
	} else {
		token, _ := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
			return t, nil
		})
		if token.Valid {
			return c.JSON(http.StatusOK, domain.RespData{
				Data: true,
			})
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, domain.RespData{
				Data: false,
			})
		}
	}
}