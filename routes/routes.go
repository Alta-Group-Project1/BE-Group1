package routes

import (
	"altaproject3/factory"
	"altaproject3/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	e.POST("users", presenter.UserPresenter.AddUser)
	e.POST("login", presenter.UserPresenter.Login)
	e.PUT("users", presenter.UserPresenter.EditData, middlewares.JWTMiddleware())
	e.GET("users", presenter.UserPresenter.GetUser, middlewares.JWTMiddleware())
	e.DELETE("users", presenter.UserPresenter.DeleteDataUser, middlewares.JWTMiddleware())
	// Events
	e.GET("events", presenter.EventPresenter.GetAllEvent)
	e.GET("events/:idEvent", presenter.EventPresenter.GetDetailEvent)
	e.GET("events/mylists", presenter.EventPresenter.GeEventOwnByUser, middlewares.JWTMiddleware())
	e.POST("events", presenter.EventPresenter.InsertNewEvent, middlewares.JWTMiddleware())
	e.DELETE("events/:idEvent", presenter.EventPresenter.DeleteEvent, middlewares.JWTMiddleware())
	e.PUT("events/:idEvent", presenter.EventPresenter.UpdateEvent, middlewares.JWTMiddleware())

	// Comments
	e.POST("/comments", presenter.CommentPresenter.AddComment, middlewares.JWTMiddleware())
	e.GET("/comments/:idEvent", presenter.UserPresenter.Login)

	// Attendees
	e.DELETE("attendees/:idAttendee", presenter.AttendeePresenter.DeleteDataAttendee, middlewares.JWTMiddleware())
	e.POST("/attendees/events/:idEvent", presenter.AttendeePresenter.InsertAttendee, middlewares.JWTMiddleware())
	e.GET("/attendees/event/:idEvent", presenter.AttendeePresenter.GetAttendeeIdEvent, middlewares.JWTMiddleware())
	e.GET("/attendees/users", presenter.AttendeePresenter.GetAttendeeIdUser, middlewares.JWTMiddleware())

	return e
}
