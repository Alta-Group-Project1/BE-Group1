package request

import (
	attendees "altaproject3/features/attendees"
)

type Attendee struct {
	UserID  int `json:"user_id" form:"user_id"`
	EventID int `json:"event_id" form:"event_id"`
}

func ToCore(req Attendee) attendees.Core {
	return attendees.Core{
		User: attendees.User{
			ID: req.UserID,
		},
		Event: attendees.Event{
			ID: req.EventID,
		},
	}
}
