package business

import (
	attendees "altaproject3/features/attendees"
	"fmt"
)

type attendeeUsecase struct {
	attendeeData attendees.Data
}

func NewAttendeeBusiness(atData attendees.Data) attendees.Business {
	return &attendeeUsecase{
		attendeeData: atData,
	}
}

func (uc *attendeeUsecase) InsertAttendee(input attendees.Core) (row int, err error) {
	check, _ := uc.attendeeData.CheckAttend(input.User.ID, input.Event.ID)
	if check == 1 {
		return -1, fmt.Errorf("already join this event")
	} else {
		row, err = uc.attendeeData.PostAttendee(input)
		return row, err
	}
}

func (uc *attendeeUsecase) DeleteAttendee(id, idUser int) (row int, err error) {
	rowDel, err := uc.attendeeData.DeleteAttendee(id, idUser)
	return rowDel, err
}

func (uc *attendeeUsecase) GetAttendeeByIdEvent(idEvent int) (resp []attendees.Core, err error) {
	resp, err = uc.attendeeData.GetAttendeeByIdEvent(idEvent)
	return resp, err
}

func (uc *attendeeUsecase) GetAttendeeByIdUser(idUser int) (resp []attendees.Core, err error) {
	resp, err = uc.attendeeData.GetAttendeeByIdUser(idUser)
	return resp, err
}
