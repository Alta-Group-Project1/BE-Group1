package presentation

import (
	"altaproject3/features/users"
	_requestUser "altaproject3/features/users/presentation/request"
	_responseUser "altaproject3/features/users/presentation/response"
	_helper "altaproject3/helper"
	"altaproject3/middlewares"
	"fmt"
	"os"

	"io"
	"net/http"
	"net/url"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
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
	id := c.Param("id")
	idUser, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to recognized ID"))
	}
	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("unauthorized"))
	}
	fullName := c.FormValue("full_name")
	userName := c.FormValue("user_name")
	password := c.FormValue("password")
	email := c.FormValue("email")
	phoneNumber := c.FormValue("phone_number")
	address := c.FormValue("address")

	var storageClient *storage.Client
	bucket := os.Getenv("DB_BUCKET")
	ctx := appengine.NewContext(c.Request())
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("misssing credentials file"))
	}
	file, err := c.FormFile("image_url")
	if err != nil {
		return err
	}
	if file.Size > 1024*1024 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("The uploaded image is too big. Please use an image less than 1MB in size"))
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	if file.Filename[len(file.Filename)-3:] != "jpg" && file.Filename[len(file.Filename)-3:] != "png" {
		if file.Filename[len(file.Filename)-4:] != "jpeg" {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("The provided file format is not allowed. Please upload a JPG or JPEG or PNG image"))
		}
	}

	sw := storageClient.Bucket(bucket).Object(file.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	if err := sw.Close(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		return err
	}
	var user = _requestUser.User{
		FullName:    fullName,
		UserName:    userName,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		Address:     address,
		ImageURL:    u.String(),
	}

	result, err := h.userBusiness.UpdateDataUser(idUser, _requestUser.ToCore(user))
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
	id := c.Param("id")
	idnya, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("id not recognize"))
	}
	if idToken != idnya {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("unauthorized"))
	}
	if errId != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to recognized ID"))
	}
	result, err := h.userBusiness.SelectUser(idnya)
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
	id := c.Param("id")
	idDel, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("id not recognize"))
	}
	if idTok != idDel {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("Unauthorized"))
	}
	_, err := h.userBusiness.DeleteUser(idDel)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete user"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("success to delete user"))
}
