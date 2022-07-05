package request

import "altaproject3/features/events"

type Event struct {
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
	UserID      int    `json:"user_id" form:"user_id"`
}

func ToCore(req Event) events.Core {
	return events.Core{
		EventName:   req.EventName,
		DateStart:   req.DateStart,
		DateFinish:  req.DateFinish,
		StartAt:     req.StartAt,
		FinishAt:    req.FinishAt,
		Price:       req.Price,
		Address:     req.Address,
		Capacity:    req.Capacity,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		User: events.User{
			ID: req.UserID,
		},
	}
}
