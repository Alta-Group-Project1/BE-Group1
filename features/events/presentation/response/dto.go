package response

import (
	"altaproject3/features/events"
	"time"
)

type Event struct {
	ID          int       `json:"id" form:"id"`
	EventName   string    `json:"event_name" form:"event_name"`
	DateStart   string    `json:"date_start" form:"date_start"`
	DateFinish  string    `json:"date_finish" form:"date_finish"`
	StartAt     string    `json:"start_at" form:"start_start"`
	FinishAt    string    `json:"finish_at" form:"finish_at"`
	Price       int       `json:"price" form:"price"`
	Address     string    `json:"address" form:"address"`
	Capacity    int       `json:"capacity" form:"capacity"`
	Description string    `json:"description" form:"description"`
	ImageURL    string    `json:"image_url" form:"image_url"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	User        User      `json:"user" form:"user"`
}

type User struct {
	ID       int    `json:"id" form:"id"`
	FullName string `json:"full_name" form:"full_name"`
}

func FromCore(data events.Core) Event {
	return Event{
		ID:          data.ID,
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
		User: User{
			ID:       data.User.ID,
			FullName: data.User.FullName,
		},
	}
}

func FromCoreList(data []events.Core) []Event {
	result := []Event{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
