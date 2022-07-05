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
	e.DELETE("users/:id", presenter.UserPresenter.DeleteDataaUser, middlewares.JWTMiddleware())

	e.GET("events", presenter.EventPresenter.GetAllEvent)

	return e
}
