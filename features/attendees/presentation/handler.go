package presentation

import (
	"altaproject3/features/attendees"
	_requestAttendees "altaproject3/features/attendees/presentation/request"
	_responseAttendees "altaproject3/features/attendees/presentation/response"
	_helper "altaproject3/helper"
	"altaproject3/middlewares"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AttendeeHandler struct {
	attendeeBusiness attendees.Business
}

func NewAttendeeHandler(business attendees.Business) *AttendeeHandler {
	return &AttendeeHandler{
		attendeeBusiness: business,
	}
}

func (h *AttendeeHandler) InsertAttendee(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("unauthorized"))
	}
	id := c.Param("idEvent")
	idEventInt, _ := strconv.Atoi(id)

	var dataAttendee = _requestAttendees.Attendee{
		UserID:  idToken,
		EventID: idEventInt,
	}
	errBind := c.Bind(&dataAttendee)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to attendee"))
	}
	row, err := h.attendeeBusiness.InsertAttendee(_requestAttendees.ToCore(dataAttendee))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed insert attendee"))
	}
	if row == 0 {
		c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert attendee"))

	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("Success to insert Attendee"))
}

func (h *AttendeeHandler) DeleteDataAttendee(c echo.Context) error {
	idTkn, errTkn := middlewares.ExtractToken(c)
	if errTkn != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}
	id := c.Param("idEvent")
	idEvent, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("id not attendee recognize"))
	}
	if idTkn == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("Unauthorized"))
	}
	_, err := h.attendeeBusiness.DeleteAttendee(idEvent, idTkn)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete attendee"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("success to delete attendee"))
}

func (h *AttendeeHandler) GetAttendeeIdEvent(c echo.Context) error {
	idTk, errTk := middlewares.ExtractToken(c)
	if errTk != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}
	id := c.Param("idEvent")
	idEvnt, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("id event recognize"))
	}
	if idTk == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("Unauthorized"))
	}
	result, err := h.attendeeBusiness.GetAttendeeByIdEvent(idEvnt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get Attendee"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("success to get data", _responseAttendees.FromCoreList(result)))
}

func (h *AttendeeHandler) GetAttendeeIdUser(c echo.Context) error {
	idTk, errTk := middlewares.ExtractToken(c)
	if errTk != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}
	if idTk == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("Unauthorized"))
	}
	result, err := h.attendeeBusiness.GetAttendeeByIdUser(idTk)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get Attendee"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("success to get data", _responseAttendees.FromCoreList(result)))
}
