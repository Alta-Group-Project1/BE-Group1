package presentation

import (
	"altaproject3/features/users"
	_requestUser "altaproject3/features/users/presentation/request"
	_responseUser "altaproject3/features/users/presentation/response"
	_helper "altaproject3/helper"
	"altaproject3/middlewares"
	"fmt"

	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: business,
	}
}

func (h *UserHandler) AddUser(c echo.Context) error {
	// var newUser _requestUser.User
	// errB := c.Bind(&newUser)
	// if errB != nil {
	// 	return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("Check your input"))
	// }
	fullName := c.FormValue("full_name")
	userName := c.FormValue("user_name")
	password := c.FormValue("password")
	email := c.FormValue("email")
	phoneNumber := c.FormValue("phone_number")
	address := c.FormValue("address")

	link, report, err := _helper.AddImageUser(c)
	if err != nil {
		return c.JSON(report["code"].(int), _helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}

	var newUser = _requestUser.User{
		FullName:    fullName,
		UserName:    userName,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		Address:     address,
		ImageURL:    link,
	}

	dataUser := _requestUser.ToCore(newUser)
	row, err := h.userBusiness.InsertUser(dataUser)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("all input must be filled"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert data"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("Succes to insert data"))
}

func (h *UserHandler) Login(c echo.Context) error {
	var userLogin _requestUser.User
	errLog := c.Bind(&userLogin)
	if errLog != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("Email or Password incorrect"))
	}
	token, userName, e := h.userBusiness.LoginUser(userLogin.Email, userLogin.Password)
	if e != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email or password incorrect"))
	}
	data := map[string]interface{}{
		"user_name": userName,
		"token":     token,
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("LOGIN SUCCES", data))
}

func (h *UserHandler) EditData(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}
	// id := c.Param("id")
	// idUser, errId := strconv.Atoi(id)
	// if errId != nil {
	// 	return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to recognized ID"))
	// }
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("unauthorized"))
	}
	fullName := c.FormValue("full_name")
	userName := c.FormValue("user_name")
	password := c.FormValue("password")
	email := c.FormValue("email")
	phoneNumber := c.FormValue("phone_number")
	address := c.FormValue("address")

	link, report, err := _helper.AddImageUser(c)
	if err != nil {
		return c.JSON(report["code"].(int), _helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	var user = _requestUser.User{
		FullName:    fullName,
		UserName:    userName,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		Address:     address,
		ImageURL:    link,
	}

	result, err := h.userBusiness.UpdateDataUser(idToken, _requestUser.ToCore(user))
	if err != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data"))
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to update data"))
	}
	if result == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("success update data"))
}

func (h *UserHandler) GetUser(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}
	// id := c.Param("id")
	// idnya, errId := strconv.Atoi(id)
	// if errId != nil {
	// 	return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("id not recognize"))
	// }
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("unauthorized"))
	}
	result, err := h.userBusiness.SelectUser(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get data user"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("success", _responseUser.FromCore(result)))
}

func (h *UserHandler) DeleteDataUser(c echo.Context) error {
	idTok, errDel := middlewares.ExtractToken(c)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}
	// id := c.Param("id")
	// idDel, errId := strconv.Atoi(id)
	// if errId != nil {
	// 	return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("id not recognize"))
	// }
	if idTok == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("Unauthorized"))
	}
	_, err := h.userBusiness.DeleteUser(idTok)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete user"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("success to delete user"))
}
