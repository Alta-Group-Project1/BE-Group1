package response

import (
	attendees "altaproject3/features/attendees"
)

type Attendee struct {
	ID int `json:"id" form:"id"`
	// CreatedAt time.Time `json:"created_at" form:"created_at"`
	User  User  `json:"user" form:"user"`
	Event Event `json:"event" form:"event"`
}

type User struct {
	ID       int    `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
	ImageURL string `json:"image_url" form:"image_url"`
}

type Event struct {
	ID        int    `json:"id" form:"id"`
	EventName string `json:"event_name" form:"event_name"`
}

func FromCore(data attendees.Core) Attendee {
	return Attendee{
		ID: data.ID,
		User: User{
			ID:       data.User.ID,
			UserName: data.User.UserName,
			ImageURL: data.User.ImageURL,
		},
		Event: Event{
			ID:        data.Event.ID,
			EventName: data.Event.EventName,
		},
	}
}

func FromCoreList(data []attendees.Core) []Attendee {
	result := []Attendee{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
