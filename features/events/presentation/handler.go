package presentation

import (
	"altaproject3/features/events"
	_requestEvent "altaproject3/features/events/presentation/request"
	_responseEvent "altaproject3/features/events/presentation/response"
	helper "altaproject3/helper"
	"altaproject3/middlewares"
	"fmt"
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
	result, totalPage, err := h.eventBusiness.GetAllEvent(limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	var resp = map[string]interface{}{
		"data":       _responseEvent.FromCoreList(result),
		"total_page": totalPage,
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccesWithData("success to get all data", resp))
}

func (h *EventHandler) GetDetailEvent(c echo.Context) error {
	idEvent := c.Param("idEvent")
	idEventInt, errId := strconv.Atoi(idEvent)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to recognize id"))
	}
	result, err := h.eventBusiness.GetDetailEvent(idEventInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccesWithData("success to get data", _responseEvent.FromCore(result)))
}

func (h *EventHandler) InsertNewEvent(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
	}

	eventName := c.FormValue("event_name")
	dateStart := c.FormValue("date_start")
	dateFinish := c.FormValue("date_finish")
	startAt := c.FormValue("start_at")
	finishAt := c.FormValue("finish_at")
	price := c.FormValue("price")
	priceInt, _ := strconv.Atoi(price)
	address := c.FormValue("address")
	description := c.FormValue("description")
	capacity := c.FormValue("capacity")
	capacityInt, _ := strconv.Atoi(capacity)

	url, report, err := helper.AddImageEvent(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}

	var newEvent = _requestEvent.Event{
		EventName:   eventName,
		DateStart:   dateStart,
		DateFinish:  dateFinish,
		StartAt:     startAt,
		FinishAt:    finishAt,
		Price:       priceInt,
		Address:     address,
		Description: description,
		UserID:      idToken,
		Capacity:    capacityInt,
		ImageURL:    url,
	}

	dataEvent := _requestEvent.ToCore(newEvent)
	row, err := h.eventBusiness.InsertNewEvent(dataEvent)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to insert event"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccesNoData("success to insert event"))
}

func (h *EventHandler) DeleteEvent(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	idEvent := c.Param("idEvent")
	idEventInt, errId := strconv.Atoi(idEvent)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to recognize id"))
	}
	result, err := h.eventBusiness.DeleteEvent(idEventInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to delete event"))
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to delete event"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccesNoData("success to delete event"))
}

func (h *EventHandler) UpdateEvent(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
	}

	idEvent := c.Param("idEvent")
	idEventInt, errId := strconv.Atoi(idEvent)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to recognize id"))
	}

	eventName := c.FormValue("event_name")
	dateStart := c.FormValue("date_start")
	dateFinish := c.FormValue("date_finish")
	startAt := c.FormValue("start_at")
	finishAt := c.FormValue("finish_at")
	price := c.FormValue("price")
	priceInt, _ := strconv.Atoi(price)
	address := c.FormValue("address")
	description := c.FormValue("description")
	capacity := c.FormValue("capacity")
	capacityInt, _ := strconv.Atoi(capacity)

	url, report, err := helper.AddImageEvent(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}

	var updateEvent = _requestEvent.Event{
		EventName:   eventName,
		DateStart:   dateStart,
		DateFinish:  dateFinish,
		StartAt:     startAt,
		FinishAt:    finishAt,
		Price:       priceInt,
		Address:     address,
		Description: description,
		UserID:      idToken,
		Capacity:    capacityInt,
		ImageURL:    url,
	}

	result, err := h.eventBusiness.UpdateEvent(idEventInt, _requestEvent.ToCore(updateEvent))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to update event"))
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to update event"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccesNoData("success to update event"))
}

func (h *EventHandler) GeEventOwnByUser(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
	}
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	result, err := h.eventBusiness.GetEventOwnByUser(idToken, limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccesWithData("success to get all data", _responseEvent.FromCoreList(result)))
}
