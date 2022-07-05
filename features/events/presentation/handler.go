package presentation

import (
	"altaproject3/features/events"
	_responseEvent "altaproject3/features/events/presentation/response"
	helper "altaproject3/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	eventBusiness events.Business
}

func NewEventHandler(business events.Business) *EventHandler {
	return &EventHandler{
		eventBusiness: business,
	}
}

func (h *EventHandler) GetAllEvent(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	result, err := h.eventBusiness.GetAllEvent(uint(limitint), uint(offsetint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccesWithData("success to get all data", _responseEvent.FromCoreList(result)))
}
