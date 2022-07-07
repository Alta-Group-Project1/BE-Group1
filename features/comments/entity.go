package comments

import "time"

type Core struct {
	ID        int
	Content   string
	User      User
	Event     Event
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID       int
	UserName string
	FullName string
	ImageURL string
}

type Event struct {
	ID         int
	EventName  string
	Price      int
	Address    string
	DateStart  string
	DateFinish string
}

type Business interface {
	InsertComment(Core) (int, error)
	GetAllComment(idEvent, limit, offset int) (data []Core, err error)
}

type Data interface {
	InsertComment(Core) (int, error)
	DataGetAllComment(idEvent, limit, offset int) (data []Core, err error)
}
