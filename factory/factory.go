package factory

import (
	_commentBusiness "altaproject3/features/comments/business"
	_commentData "altaproject3/features/comments/data"
	_commentPresentation "altaproject3/features/comments/presentation"
	_eventBusiness "altaproject3/features/events/business"
	_eventData "altaproject3/features/events/data"
	_eventPresentation "altaproject3/features/events/presentation"
	_userBusiness "altaproject3/features/users/business"
	_userData "altaproject3/features/users/data"
	_userPresentation "altaproject3/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter    *_userPresentation.UserHandler
	EventPresenter   *_eventPresentation.EventHandler
	CommentPresenter *_commentPresentation.CommentHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	UserPresentation := _userPresentation.NewUserHandler(userBusiness)
	eventData := _eventData.NewEventRepository(dbConn)
	eventBusiness := _eventBusiness.NewEventBusiness(eventData)
	eventPresentation := _eventPresentation.NewEventHandler(eventBusiness)
	commentData := _commentData.NewCommentRepository(dbConn)
	commentBusiness := _commentBusiness.NewCommentBusiness(commentData)
	commentPresentation := _commentPresentation.NewCommentHandler(commentBusiness)
	return Presenter{
		UserPresenter:    UserPresentation,
		EventPresenter:   eventPresentation,
		CommentPresenter: commentPresentation,
	}
}
