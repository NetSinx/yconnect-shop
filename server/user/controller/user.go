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
	"github.com/NetSinx/yconnect-shop/server/user/utils"
	"github.com/golang-jwt/jwt/v5"
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

func (u userController) LoginUser(c echo.Context) error {
	var userLogin domain.UserLogin

	if err := c.Bind(&userLogin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	accessToken, refreshToken, user_id, err := u.userService.LoginUser(userLogin)
	if err != nil && err.Error() == "email atau password salah" {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: "Email atau password Anda salah.",
		})
	} else if err != nil && err.Error() == echo.ErrBadRequest.Error() {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: err.Error(),
		})
	}

	utils.SetCookies(c, "user_session", accessToken, time.Now().Add(15 * time.Minute))
	utils.SetCookies(c, "refresh_token", refreshToken, time.Now().Add(1 * time.Hour))
	utils.SetCookies(c, "user_id", user_id, time.Now().Add(15 * time.Minute))
	utils.SetCookies(c, "tz", time.Now().String(), time.Now().Add(15 * time.Minute))

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "User berhasil login",
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

	getDbUser, _ := u.userService.GetUser(user, username)
	
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

	username := c.Param("username")

	findUser, err := u.userService.GetUser(users, username)
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

	err := u.userService.DeleteUser(users, username)
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

func (u userController) UserLogout(c echo.Context) error {
	session, err := c.Cookie("user_session")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: "user session in cookie is not available",
		})
	}
	session.Path = "/"
	session.MaxAge = -1

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: "refresh token in cookie is not available",
		})
	}
	refreshToken.Path = "/"
	refreshToken.MaxAge = -1

	user_id, err := c.Cookie("user_id")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: "user id in cookie is not available",
		})
	}
	user_id.Path = "/"
	user_id.MaxAge = -1

	tz, err := c.Cookie("tz")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: "timezone in cookie is not available",
		})
	}
	tz.Path = "/"
	tz.MaxAge = -1

	c.SetCookie(session)
	c.SetCookie(refreshToken)
	c.SetCookie(user_id)
	c.SetCookie(tz)

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "User berhasil logout",
	})
}

func (u userController) GetAccessToken(c echo.Context) error {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: "refresh token not available",
		})
	}

	token, err := jwt.ParseWithClaims(refreshToken.Value, &utils.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("adminyasinnetsinx_15"), nil
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: err.Error(),
		})
	}
	if !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: "your token is invalid",
		})
	}

	claims := token.Claims.(*utils.CustomClaims)
	if claims.Username == "" && claims.Role == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: "your claims is invalid",
		})
	}

	newAccessToken := utils.GenerateAccessToken(claims.Username, claims.Role)
	utils.SetCookies(c, "user_session", newAccessToken, time.Now().Add(15 * time.Minute))
	utils.SetCookies(c, "user_id", claims.Username, time.Now().Add(15 * time.Minute))
	utils.SetCookies(c, "tz", time.Now().String(), time.Now().Add(15 * time.Minute))

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Access token regenerated successfully",
	})
}
