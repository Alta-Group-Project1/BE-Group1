package attendees

import "time"

type Core struct {
	ID        int
	User      User
	Event     Event
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID       int
	UserName string
}

type Event struct {
	ID        int
	EventName string
}

type Business interface {
	InsertAttendee(data Core) (row int, err error)
}

type Data interface {
	PostAttendee(data Core) (row int, err error)
}
