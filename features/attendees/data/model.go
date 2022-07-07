package data

import (
	attendees "altaproject3/features/attendees"

	"gorm.io/gorm"
)

type Attendee struct {
	gorm.Model
	UserID  uint `json:"user_id" form:"user_id"`
	EventID uint `json:"event_id" form:"event_id"`
	User    User
	Event   Event
}

type User struct {
	gorm.Model
	UserName string `json:"user_name" form:"user_name"`
	ImageURL string `json:"image_url" form:"image_url"`
	Attendee []Attendee
}

type Event struct {
	gorm.Model
	EventName string `json:"event_name" form:"event_name"`
	Attendee  []Attendee
}

func (data *Attendee) toCore() attendees.Core {
	return attendees.Core{
		ID: int(data.ID),
		User: attendees.User{
			ID:       int(data.User.ID),
			UserName: data.User.UserName,
			ImageURL: data.User.ImageURL,
		},
		Event: attendees.Event{
			ID:        int(data.Event.ID),
			EventName: data.Event.EventName,
		},
	}
}

func toCoreList(data []Attendee) []attendees.Core {
	result := []attendees.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core attendees.Core) Attendee {
	return Attendee{
		UserID:  uint(core.User.ID),
		EventID: uint(core.Event.ID),
	}
}
