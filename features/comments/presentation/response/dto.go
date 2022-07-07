package response

import (
	"altaproject3/features/comments"
	"time"
)

type Comment struct {
	ID        int       `json:"id" form:"id"`
	Content   string    `json:"content" form:"content"`
	User      User      `json:"user" form:"user"`
	Event     Event     `json:"event" form:"event"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

type User struct {
	ID       int    `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
	FullName string `json:"full_name" form:"full_name"`
	ImageURL string `json:"image_url" form:"image_url"`
}

type Event struct {
	ID int `json:"id" form:"id"`
}

func FromCore(data comments.Core) Comment {
	return Comment{
		ID:        data.ID,
		Content:   data.Content,
		CreatedAt: data.CreatedAt,
		User: User{
			ID:       data.User.ID,
			UserName: data.User.UserName,
			FullName: data.User.FullName,
			ImageURL: data.User.ImageURL,
		},
		Event: Event{
			ID: data.Event.ID,
		},
	}
}

func FromCoreList(data []comments.Core) []Comment {
	result := []Comment{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
