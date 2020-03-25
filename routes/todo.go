package routes

import (
	"github.com/jorgeAM/echo-api/controllers"
	"github.com/labstack/echo/v4"
)

func initTodoRoutes(e *echo.Echo) {
	e.GET("/", controllers.HiTodo)
}
