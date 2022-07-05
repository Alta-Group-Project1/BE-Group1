package business

import "altaproject3/features/events"

type eventUsecase struct {
	eventData events.Data
}

func NewEventBusiness(evData events.Data) events.Business {
	return &eventUsecase{
		eventData: evData,
	}
}

func (uc *eventUsecase) GetAllEvent(limit, offset uint) (resp []events.Core, err error) {
	resp, err = uc.eventData.SelectAllEvent(limit, offset)
	return resp, err
}
