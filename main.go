package main

import (
	"altaproject3/config"
	"altaproject3/factory"
	"altaproject3/routes"

	"altaproject3/middlewares"
)

func main() {
	dbConn := config.InitDB()

	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)
	middlewares.LogMiddleware(*e)
	e.Logger.Fatal(e.Start(":5000"))
}
