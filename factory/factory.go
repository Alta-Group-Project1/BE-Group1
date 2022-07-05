package factory

import (
	_eventBusiness "altaproject3/features/events/business"
	_eventData "altaproject3/features/events/data"
	_eventPresentation "altaproject3/features/events/presentation"
	_userBusiness "altaproject3/features/users/business"
	_userData "altaproject3/features/users/data"
	_userPresentation "altaproject3/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter  *_userPresentation.UserHandler
	EventPresenter *_eventPresentation.EventHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	UserPresentation := _userPresentation.NewUserHandler(userBusiness)
	eventData := _eventData.NewEventRepository(dbConn)
	eventBusiness := _eventBusiness.NewEventBusiness(eventData)
	eventPresentation := _eventPresentation.NewEventHandler(eventBusiness)

	return Presenter{
		UserPresenter:  UserPresentation,
		EventPresenter: eventPresentation,
	}
}
