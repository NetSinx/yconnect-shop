package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"github.com/NetSinx/yconnect-shop/user/app/model"
	"github.com/NetSinx/yconnect-shop/user/service"
	"github.com/NetSinx/yconnect-shop/user/utils"
	validation "github.com/go-passwd/validator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	var users *model.User
	
	if err := c.Bind(&users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: "Request data doesn't match!",
		})
	}

	if err := validator.New().Struct(users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: err.Error(),
		})
	}
	
	passwordValidation := validation.New(validation.MinLength(5, errors.New("too short")), validation.ContainsAtLeast("!@#$%&*", 1, errors.New("password must be contain symbol")), validation.ContainsAtLeast("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1, errors.New("password must be contain uppercase character")), validation.ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 2, errors.New("password must be contain lowercase character")), validation.ContainsAtLeast("0123456789", 1, errors.New("password must be contain number")))

	if err := passwordValidation.Validate(users.Password); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: err.Error(),
		})
	}
	
	passwdHash, _ := bcrypt.GenerateFromPassword([]byte(users.Password), 15)
	
	users = &model.User{
		Name: users.Name,
		Username: users.Username,
		Email: users.Email,
		Alamat: users.Alamat,
		NoTelp: users.NoTelp,
		Product: users.Product,
		Password: string(passwdHash),
	}

	if err := u.userService.RegisterUser(*users); err != nil {
		fmt.Printf("Error message: %v", err)

		return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
			Code: http.StatusConflict,
			Status: "Data Conflict",
			Message: "Data was availabled!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessCUD{
		Code: http.StatusOK,
		Status: "OK",
		Message: "User registered successfully!",
	})
}

func (u userController) LoginUser(c echo.Context) error {
	var userLogin *model.UserLogin

	if err := c.Bind(&userLogin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: "Request data doesn't match!",
		})
	}

	if err := validator.New().Struct(userLogin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: err.Error(),
		})
	}

	passwordValidation := validation.New(validation.MinLength(5, errors.New("too short")))

	if err := passwordValidation.Validate(userLogin.Password); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: err.Error(),
		})
	}

	email := userLogin.Email

	user, err := u.userService.LoginUser(email)
	if err != nil {
		fmt.Printf("Error message: %v", err)

		return echo.NewHTTPError(http.StatusUnauthorized, utils.SuccessCUD{
			Code: http.StatusUnauthorized,
			Status: "Unauthorized",
			Message: "Your username/password is wrong!",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.SuccessCUD{
			Code: http.StatusUnauthorized,
			Status: "Unauthorized",
			Message: "Your username/password is wrong!",
		})
	}

	jwtToken := utils.JWTAuth()

	return c.JSON(http.StatusOK, map[string]string{
		"code": strconv.Itoa(http.StatusOK),
		"status": "OK",
		"token": jwtToken,
	})
}

func (u userController) ListUsers(c echo.Context) error {
	var users []model.User
	
	listUsers, err := u.userService.ListUsers(users)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: "Internal Server Error",
			Message: "Maaf, ada kesalahan pada server",
		})
	}

	for i, user := range listUsers {
		var preloadProduct utils.PreloadProducts

		responseData, err := http.Get(fmt.Sprintf("http://localhost:8000/products/user/%d", user.Id))
		if err != nil {
			return c.JSON(http.StatusOK, utils.SuccessGet{
				Code: http.StatusOK,
				Status: "OK",
				Data: listUsers,
			})
		}

		if err := json.NewDecoder(responseData.Body).Decode(&preloadProduct); err != nil {
			return err
		}

		listUsers[i].Product = append(listUsers[i].Product, preloadProduct.Data...)
	}

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: "OK",
		Data: listUsers,
	})
}

func (u userController) UpdateUser(c echo.Context) error {
	var users *model.User

	if err := c.Bind(&users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: "Request doesn't match!",
		})
	}

	if err := validator.New().Struct(users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: err.Error(),
		})
	}

	getId, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(getId)

	err := u.userService.UpdateUser(*users, id)
	if (err != nil  && err.Error() != "record not found") {
		return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
			Code: http.StatusConflict,
			Status: "Data Conflict",
			Message: "Data was existing!",
		})
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Message: "User cannot be found!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessCUD{
		Code: http.StatusOK,
		Status: "OK",
		Message: "User updated successfully!",
	})
}

func (u userController) FindUser(c echo.Context) error {
	var users model.User
	var preloadProduct utils.PreloadProducts

	getId, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(getId)

	findUser, err := u.userService.FindUser(users, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Message: "User cannot be found!",
		})
	}


	responseData, err := http.Get(fmt.Sprintf("http://localhost:8000/products/user/%d", findUser.Id))
	if err != nil {
		return c.JSON(http.StatusOK, utils.SuccessGet{
			Code: http.StatusOK,
			Status: "OK",
			Data: findUser,
		})
	}

	if err := json.NewDecoder(responseData.Body).Decode(&preloadProduct); err != nil {
		return err
	}

	findUser.Product = preloadProduct.Data

	return c.JSON(http.StatusOK, utils.SuccessGet{
		Code: http.StatusOK,
		Status: "OK",
		Data: findUser,
	})
}

func (u userController) DeleteUser(c echo.Context) error {
	var users model.User

	getId, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(getId)

	if err := u.userService.DeleteUser(users, id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Message: "User cannot be found!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessCUD{
		Code: http.StatusOK,
		Status: "OK",
		Message: "User deleted successfully!",
	})
}