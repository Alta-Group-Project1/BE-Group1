package business

import (
	"altaproject3/features/events"
	"errors"
)

type eventUsecase struct {
	eventData events.Data
}

func NewEventBusiness(evData events.Data) events.Business {
	return &eventUsecase{
		eventData: evData,
	}
}

func (uc *eventUsecase) GetAllEvent(limit, offset int) (resp []events.Core, totalPage int, err error) {
	resp, err = uc.eventData.SelectAllEvent(limit, offset)
	totalData, errCount := uc.eventData.CountEventData()
	if errCount != nil {
		totalPage = 0
	} else {
		if limit == 0 {
			limit = totalData
		}
		if totalData%limit != 0 {
			totalPage = (totalData / limit) + 1
		} else {
			totalPage = totalData / limit
		}
	}
	return resp, totalPage, err
}

func (uc *eventUsecase) GetDetailEvent(idEvent int) (resp events.Core, err error) {
	resp, err = uc.eventData.SelectEvent(idEvent)
	return resp, err
}

func (uc *eventUsecase) InsertNewEvent(input events.Core) (row int, err error) {
	if input.EventName == "" || input.DateStart == "" || input.DateFinish == "" || input.StartAt == "" || input.FinishAt == "" || input.User.ID == 0 || input.Price == 0 || input.Capacity == 0 || input.Description == "" {
		return -1, errors.New("all input data must be filled")
	}
	row, err = uc.eventData.InsertEvent(input)
	return row, err
}

func (uc *eventUsecase) DeleteEvent(idEvent int) (row int, err error) {
	row, err = uc.eventData.DeleteEvent(idEvent)
	return row, err
}

func (uc *eventUsecase) UpdateEvent(idEvent int, input events.Core) (row int, err error) {
	if input.EventName == "" || input.DateStart == "" || input.DateFinish == "" || input.StartAt == "" || input.FinishAt == "" || input.User.ID == 0 || input.Price == 0 || input.Capacity == 0 || input.Description == "" {
		return -1, errors.New("all input data must be filled")
	}
	row, err = uc.eventData.UpdateEvent(idEvent, input)
	return row, err
}

func (uc *eventUsecase) GetEventOwnByUser(idUser, limit, offset int) (resp []events.Core, err error) {
	resp, err = uc.eventData.SelectEventByUserId(idUser, limit, offset)
	return resp, err
}
