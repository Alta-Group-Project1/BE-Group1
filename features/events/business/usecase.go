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

func (uc *eventUsecase) GetAllEvent(limit, offset int) (resp []events.Core, err error) {
	resp, err = uc.eventData.SelectAllEvent(limit, offset)
	return resp, err
}

func (uc *eventUsecase) GetDetailEvent(idEvent int) (resp events.Core, err error) {
	resp, err = uc.eventData.SelectEvent(idEvent)
	return resp, err
}

func (uc *eventUsecase) InsertNewEvent(input events.Core) (row int, err error) {
	if input.EventName == "" || input.DateStart == "" || input.DateFinish == "" || input.StartAt == "" || input.FinishAt == "" || input.User.ID == 0 || input.Price == 0 || input.Capacity == 0 {
		return -1, errors.New("all input data must be filled")
	}
	row, err = uc.eventData.InsertEvent(input)
	return row, err
}
