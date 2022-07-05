package business

import (
	attendees "altaproject3/features/attendees"
	"errors"
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
	if input.User.ID == 0 || input.Event.ID == 0 {
		return -1, errors.New("all input data must be filled")
	}
	row, err = uc.attendeeData.PostAttendee(input)
	return row, err
}
