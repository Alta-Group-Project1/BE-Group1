package presentation

import (
	"altaproject3/features/comments"
	_requestComment "altaproject3/features/comments/presentation/request"
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
	id := c.Param("idEvent")
	convert, _ := strconv.Atoi(id)
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}
	var dataComment = _requestComment.Comment{
		UserID:  idToken,
		EventID: convert,
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
