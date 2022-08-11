package server

import (
	"myapp/routes"
	"myapp/routes/task"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	e.GET("/health_check", routes.Health)
	e.POST("/task/create", task.Create)

	e.Logger.Fatal(e.Start(":6000"))
}
