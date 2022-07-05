package data

import (
	"altaproject3/features/events"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	EventName   string `json:"event_name" form:"event_name"`
	DateStart   string `json:"date_start" form:"date_start"`
	DateFinish  string `json:"date_finish" form:"date_finish"`
	StartAt     string `json:"start_at" form:"start_start"`
	FinishAt    string `json:"finish_at" form:"finish_at"`
	Price       int    `json:"price" form:"price"`
	Address     string `json:"address" form:"address"`
	Capacity    int    `json:"capacity" form:"capacity"`
	Description string `json:"description" form:"description"`
	ImageURL    string `json:"image_url" form:"image_url"`
	UserID      uint   `json:"user_id" form:"user_id"`
	User        User
}

type User struct {
	gorm.Model
	FullName string `json:"full_name" form:"full_name"`
	Event    []Event
}

func (data *Event) toCore() events.Core {
	return events.Core{
		ID:          int(data.ID),
		EventName:   data.EventName,
		DateStart:   data.DateStart,
		DateFinish:  data.DateFinish,
		StartAt:     data.StartAt,
		FinishAt:    data.FinishAt,
		Price:       data.Price,
		Address:     data.Address,
		Capacity:    data.Capacity,
		Description: data.Description,
		ImageURL:    data.ImageURL,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		User: events.User{
			ID:       int(data.User.ID),
			FullName: data.User.FullName,
		},
	}
}

func toCoreList(data []Event) []events.Core {
	result := []events.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core events.Core) Event {
	return Event{
		EventName:   core.EventName,
		DateStart:   core.DateStart,
		DateFinish:  core.DateFinish,
		StartAt:     core.StartAt,
		FinishAt:    core.FinishAt,
		Price:       core.Price,
		Address:     core.Address,
		Capacity:    core.Capacity,
		Description: core.Description,
		UserID:      uint(core.User.ID),
		ImageURL:    core.ImageURL,
	}
}
