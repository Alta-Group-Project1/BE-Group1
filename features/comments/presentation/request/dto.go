package request

import "altaproject3/features/comments"

type Comment struct {
	Content string `json:"content" form:"content"`
	UserID  int    `json:"user_id" form:"user_id"`
	EventID int    `json:"event_id" form:"event_id"`
}

func ToCore(req Comment) comments.Core {
	return comments.Core{
		Content: req.Content,
		User: comments.User{
			ID: req.UserID,
		},
		Event: comments.Event{
			ID: req.EventID,
		},
	}
}
