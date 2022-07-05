package presentation

import (
	"altaproject3/features/attendees"
	_requestAttendees "altaproject3/features/attendees/presentation/request"
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
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to insert attendee"))
	}
	row, err := h.attendeeBusiness.InsertAttendee(_requestAttendees.ToCore(dataAttendee))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert attendee"))
	}
	if row == 0 {
		c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert attendee"))

	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("Success to insert Attendee"))
}
