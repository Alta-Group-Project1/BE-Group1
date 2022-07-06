package response

import (
	attendees "altaproject3/features/attendees"
)

type Attendee struct {
	ID int `json:"id" form:"id"`
	// CreatedAt time.Time `json:"created_at" form:"created_at"`
	User  User
	Event Event
}

type User struct {
	ID       int    `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
}

type Event struct {
	ID       int    `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
}

func FromCore(data attendees.Core) Attendee {
	return Attendee{
		ID: data.ID,
		User: User{
			ID:       data.User.ID,
			UserName: data.User.UserName,
		},
		Event: Event{
			ID:       data.Event.ID,
			UserName: data.Event.EventName,
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
