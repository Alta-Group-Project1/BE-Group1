package business

import (
	attendees "altaproject3/features/attendees"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAttendeeDataFailed struct{}

func (mock mockAttendeeDataFailed) PostAttendee(data attendees.Core) (int, error) {
	return 0, fmt.Errorf("failed to insert data")
}
func (mock mockAttendeeDataFailed) DeleteAttendee(id, idUser int) (int, error) {
	return 0, fmt.Errorf("failed to delete data")
}

func (mock mockAttendeeDataFailed) GetAttendeeByIdEvent(idEvent int) ([]attendees.Core, error) {
	return nil, fmt.Errorf("failed to get data")
}

func (mock mockAttendeeDataFailed) GetAttendeeByIdUser(idUser int) ([]attendees.Core, error) {
	return nil, fmt.Errorf("failed to get data")
}

func (mock mockAttendeeDataFailed) CheckAttend(idUser, idEvent int) (int, error) {
	return 1, nil
}

type mockAttendeeData struct{}

func (mock mockAttendeeData) PostAttendee(data attendees.Core) (int, error) {
	return 1, nil
}
func (mock mockAttendeeData) DeleteAttendee(id, idUser int) (int, error) {
	return 1, nil
}

func (mock mockAttendeeData) GetAttendeeByIdEvent(idEvent int) ([]attendees.Core, error) {
	return []attendees.Core{
		{ID: 1, User: attendees.User{ID: 1, UserName: "Bambang", ImageURL: "storage.cloud.com/profile.jpg"}, Event: attendees.Event{ID: 1, EventName: "Konser Akbar"}},
	}, nil
}

func (mock mockAttendeeData) GetAttendeeByIdUser(idUser int) ([]attendees.Core, error) {
	return []attendees.Core{
		{ID: 1, User: attendees.User{ID: 1, UserName: "Bambang", ImageURL: "storage.cloud.com/profile.jpg"}, Event: attendees.Event{ID: 1, EventName: "Konser Akbar"}},
	}, nil
}

func (mock mockAttendeeData) CheckAttend(idUser, idEvent int) (int, error) {
	return -1, fmt.Errorf("")
}

func TestInsertAttendee(t *testing.T) {
	t.Run("Test Insert Attendees Success", func(t *testing.T) {
		attendeesBusiness := NewAttendeeBusiness(mockAttendeeData{})
		var data = attendees.Core{
			User: attendees.User{
				ID: 1,
			},
			Event: attendees.Event{
				ID: 1,
			},
		}
		result, err := attendeesBusiness.InsertAttendee(data)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Insert Attendees Failed", func(t *testing.T) {
		attendeesBusiness := NewAttendeeBusiness(mockAttendeeDataFailed{})
		var data = attendees.Core{
			User: attendees.User{
				ID: 1,
			},
			Event: attendees.Event{
				ID: 1,
			},
		}
		result, err := attendeesBusiness.InsertAttendee(data)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
}

func TestDeleteAttendee(t *testing.T) {
	t.Run("Test Delete Attendees Success", func(t *testing.T) {
		attendeesBusiness := NewAttendeeBusiness(mockAttendeeData{})
		result, err := attendeesBusiness.DeleteAttendee(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Delete Attendees Failed", func(t *testing.T) {
		attendeesBusiness := NewAttendeeBusiness(mockAttendeeDataFailed{})
		result, err := attendeesBusiness.DeleteAttendee(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestGetAttendeeByIdEvent(t *testing.T) {
	t.Run("Test Get Attendees By Id Event Success", func(t *testing.T) {
		attendeesBusiness := NewAttendeeBusiness(mockAttendeeData{})
		result, err := attendeesBusiness.GetAttendeeByIdEvent(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result[0].Event.ID)
	})
	t.Run("Test Get Attendees By Id Event Failed", func(t *testing.T) {
		attendeesBusiness := NewAttendeeBusiness(mockAttendeeDataFailed{})
		result, err := attendeesBusiness.GetAttendeeByIdEvent(1)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestGetAttendeeByIdUser(t *testing.T) {
	t.Run("Test Get Attendees By Id User Success", func(t *testing.T) {
		attendeesBusiness := NewAttendeeBusiness(mockAttendeeData{})
		result, err := attendeesBusiness.GetAttendeeByIdUser(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result[0].User.ID)
	})
	t.Run("Test Get Attendees By Id User Failed", func(t *testing.T) {
		attendeesBusiness := NewAttendeeBusiness(mockAttendeeDataFailed{})
		result, err := attendeesBusiness.GetAttendeeByIdUser(1)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
