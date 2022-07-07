package data

import (
	"altaproject3/features/comments"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `json:"content" form:"content"`
	UserID  int    `json:"user_id" form:"user_id"`
	EventID int    `json:"event_id" form:"event_id"`
	User    User
	Event   Event
}

type User struct {
	gorm.Model
	UserName string `json:"user_name" form:"user_name"`
	FullName string `json:"full_name" form:"full_name"`
	ImageURL string `json:"image_url" form:"image_url"`
	Comment  []Comment
}

type Event struct {
	gorm.Model
	EventName  string `json:"event_name" form:"event_name"`
	Price      int    `json:"price" form:"price"`
	Address    string `json:"address" form:"address"`
	DateStart  string `json:"date_start" form:"date_start"`
	DateFinish string `json:"date_finish" form:"date_finish"`
	Comment    []Comment
}

func (data *Comment) toCore() comments.Core {
	return comments.Core{
		ID:        int(data.ID),
		Content:   data.Content,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		User: comments.User{
			ID:       int(data.User.ID),
			UserName: data.User.UserName,
			FullName: data.User.FullName,
			ImageURL: data.User.ImageURL,
		},

		Event: comments.Event{
			ID:         int(data.Event.ID),
			EventName:  data.Event.EventName,
			Price:      data.Event.Price,
			Address:    data.Event.Address,
			DateStart:  data.Event.DateStart,
			DateFinish: data.Event.DateFinish,
		},
	}
}

func toCoreList(data []Comment) []comments.Core {
	result := []comments.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core comments.Core) Comment {
	return Comment{
		Content: core.Content,
		UserID:  core.User.ID,
		EventID: core.Event.ID,
	}
}
