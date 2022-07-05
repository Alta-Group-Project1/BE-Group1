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
	GetAllEvent(limit, offset uint) (data []Core, err error)
}

type Data interface {
	SelectAllEvent(limit, offset uint) (data []Core, err error)
}
