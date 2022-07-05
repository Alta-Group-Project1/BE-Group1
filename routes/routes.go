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
	e.PUT("users/:id", presenter.UserPresenter.EditData, middlewares.JWTMiddleware())
	e.GET("users/:id", presenter.UserPresenter.GetUser, middlewares.JWTMiddleware())
	e.DELETE("users/:id", presenter.UserPresenter.DeleteDataUser, middlewares.JWTMiddleware())

	e.GET("events", presenter.EventPresenter.GetAllEvent)
	e.GET("events/:idEvent", presenter.EventPresenter.GetDetailEvent)
	e.POST("events", presenter.EventPresenter.InsertNewEvent, middlewares.JWTMiddleware())
	e.DELETE("events/:idEvent", presenter.EventPresenter.DeleteEvent, middlewares.JWTMiddleware())
	e.PUT("events/:idEvent", presenter.EventPresenter.UpdateEvent, middlewares.JWTMiddleware())

	// Comments
	e.POST("/comments/:idEvent", presenter.CommentPresenter.AddComment, middlewares.JWTMiddleware())
	e.GET("/comments/:idEvent", presenter.UserPresenter.Login)

	// Attendees
	e.POST("/attendees/:idEvent", presenter.AttendeePresenter.InsertAttendee)

	return e
}
