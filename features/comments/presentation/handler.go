package presentation

import (
	"altaproject3/features/comments"
	_requestComment "altaproject3/features/comments/presentation/request"
	_responseComment "altaproject3/features/comments/presentation/response"
	_helper "altaproject3/helper"
	"altaproject3/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentBusiness comments.Business
}

func NewCommentHandler(business comments.Business) *CommentHandler {
	return &CommentHandler{
		commentBusiness: business,
	}
}

func (h *CommentHandler) AddComment(c echo.Context) error {

	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}
	var dataComment = _requestComment.Comment{
		UserID: idToken,
	}
	errBind := c.Bind(&dataComment)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, _helper.ResponseFailed("Failed to Bind data"))
	}
	var row, err = h.commentBusiness.InsertComment(_requestComment.ToCore(dataComment))
	if err != nil {
		c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert comment"))

	}
	if row == 0 {
		c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert comment"))

	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("succes to insert comment"))
}

func (h *CommentHandler) GetAllComment(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)

	if errToken != nil {
		c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("Unauthorized"))
	}

	idEvent := c.Param("idEvent")
	idEventint, _ := strconv.Atoi(idEvent)
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	result, err := h.commentBusiness.GetAllComment(idEventint, limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("success to get all data", _responseComment.FromCoreList(result)))
}
