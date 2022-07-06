package events

import "time"

type Core struct {
	ID          int
	EventName   string
	DateStart   string
	DateFinish  string
	StartAt     string
	FinishAt    string
	Price       int
	Address     string
	Capacity    int
	Description string
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User
}

type User struct {
	ID       int
	FullName string
}

type Business interface {
	GetAllEvent(limit, offset int) (data []Core, err error)
	GetDetailEvent(idEvent int) (data Core, err error)
	GeEventOwnByUser(idUser, limit, offset int) (data []Core, err error)
	InsertNewEvent(data Core) (row int, err error)
	DeleteEvent(idEvent int) (row int, err error)
	UpdateEvent(idEvent int, data Core) (row int, err error)
}

type Data interface {
	SelectAllEvent(limit, offset int) (data []Core, err error)
	SelectEvent(idEvent int) (data Core, err error)
	SelectEventByUserId(idUser, limit, offset int) (data []Core, err error)
	InsertEvent(data Core) (row int, err error)
	DeleteEvent(idEvent int) (row int, err error)
	UpdateEvent(idEvent int, data Core) (row int, err error)
}
