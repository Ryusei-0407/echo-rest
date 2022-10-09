package server

import (
	"app/routes"
	"app/routes/task"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	e.GET("/health_check", routes.Health)
	e.POST("/api/v1/tasks/create", task.Create)

	e.Logger.Fatal(e.Start(":6000"))
}
