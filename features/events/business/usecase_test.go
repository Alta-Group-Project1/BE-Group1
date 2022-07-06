package business

import (
	"altaproject3/features/events"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockEventData struct{}

func (mock mockEventData) SelectAllEvent(limit, offset int) ([]events.Core, error) {
	return []events.Core{
		{ID: 1, EventName: "Konser Berbakti", DateStart: "19/12/2022", DateFinish: "19/12/2022", StartAt: "15.00", FinishAt: "22.00", Price: 15000, Address: "Lapangan Mandala Krida", Capacity: 1000, Description: "Konser bertajuk kebersamaan", ImageURL: "www.storage.cloud.com", User: events.User{ID: 1, FullName: "Andi Setiawan"}},
	}, nil
}

func (mock mockEventData) SelectEvent(idEvent int) (events.Core, error) {
	return events.Core{
		ID: 1, EventName: "Konser Berbakti", DateStart: "19/12/2022", DateFinish: "19/12/2022", StartAt: "15.00", FinishAt: "22.00", Price: 15000, Address: "Lapangan Mandala Krida", Capacity: 1000, Description: "Konser bertajuk kebersamaan", ImageURL: "www.storage.cloud.com",
	}, nil
}

func (mock mockEventData) InsertEvent(data events.Core) (int, error) {
	return 1, nil
}

func (mock mockEventData) DeleteEvent(idEvent int) (int, error) {
	return 1, nil
}

func (mock mockEventData) UpdateEvent(idEvent int, data events.Core) (int, error) {
	return 1, nil
}

func (mock mockEventData) SelectEventByUserId(idUser, limit, offset int) ([]events.Core, error) {
	return []events.Core{
		{ID: 1, EventName: "Konser Berbakti", DateStart: "19/12/2022", DateFinish: "19/12/2022", StartAt: "15.00", FinishAt: "22.00", Price: 15000, Address: "Lapangan Mandala Krida", Capacity: 1000, Description: "Konser bertajuk kebersamaan", ImageURL: "www.storage.cloud.com", User: events.User{ID: 1, FullName: "Andi Setiawan"}},
	}, nil
}

func (mock mockEventData) CountEventData() (count int, err error) {
	return 20, nil
}

type mockEventDataFailed struct{}

func (mock mockEventDataFailed) SelectAllEvent(limit, offset int) ([]events.Core, error) {
	return nil, fmt.Errorf("Failed to select data")
}

func (mock mockEventDataFailed) SelectEvent(idEvent int) (events.Core, error) {
	return events.Core{}, fmt.Errorf("Failed to select data")
}

func (mock mockEventDataFailed) InsertEvent(data events.Core) (int, error) {
	return 0, fmt.Errorf("Failed to insert data")
}

func (mock mockEventDataFailed) DeleteEvent(idEvent int) (int, error) {
	return 0, fmt.Errorf("Failed to delete data")
}

func (mock mockEventDataFailed) UpdateEvent(idEvent int, data events.Core) (int, error) {
	return 0, fmt.Errorf("Failed to update data")
}

func (mock mockEventDataFailed) SelectEventByUserId(idUser, limit, offset int) ([]events.Core, error) {
	return nil, fmt.Errorf("Failed to select data")
}

func (mock mockEventDataFailed) CountEventData() (int, error) {
	return 0, fmt.Errorf("Failed to count data")
}

func TestGetAllEvent(t *testing.T) {
	t.Run("Test Get All Event Success", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventData{})
		result, totalPage, err := eventBusiness.GetAllEvent(0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, result[0].ID)
		assert.Equal(t, 1, totalPage)
	})
	t.Run("Test Get All Event Success if limit value not zero", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventData{})
		result, totalPage, err := eventBusiness.GetAllEvent(5, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, result[0].ID)
		assert.Equal(t, 4, totalPage)
	})
	t.Run("Test Get All Event Success if the modulo is not zero", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventData{})
		result, totalPage, err := eventBusiness.GetAllEvent(3, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, result[0].ID)
		assert.Equal(t, 7, totalPage)
	})
	t.Run("Test Get All Event Failed", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		result, totalPage, err := eventBusiness.GetAllEvent(0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, 0, totalPage)
	})
}

func TestGetDetailEvent(t *testing.T) {
	t.Run("Test Get Detail Event Success", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventData{})
		result, err := eventBusiness.GetDetailEvent(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result.ID)
	})
	t.Run("Test Get Detail Event Failed", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		result, err := eventBusiness.GetDetailEvent(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result.ID)
	})
}
