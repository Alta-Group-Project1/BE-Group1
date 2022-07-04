package presentation

import (
	"altaproject3/features/users"
	_requestUser "altaproject3/features/users/presentation/request"
	_helper "altaproject3/helper"
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
	var newUser _requestUser.User
	errB := c.Bind(&newUser)
	if errB != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("Check your input"))
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
