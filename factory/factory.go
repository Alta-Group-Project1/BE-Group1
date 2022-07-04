package factory

import (
	_userBusiness "altaproject3/features/users/business"
	_userData "altaproject3/features/users/data"
	_userPresentation "altaproject3/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter *_userPresentation.UserHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	UserPresentation := _userPresentation.NewUserHandler(userBusiness)

	return Presenter{
		UserPresenter: UserPresentation,
	}
}
