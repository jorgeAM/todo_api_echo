package routes

import (
	"github.com/jorgeAM/echo-api/controllers"
	"github.com/labstack/echo/v4"
)

func initTodoRoutes(e *echo.Echo) {
	g := e.Group("todos")
	g.GET("", controllers.HiTodo)
	g.POST("", controllers.NewTodo)
	g.GET("/:id", controllers.GetTodo)
	g.PUT("/:id", controllers.UpdateTodo)
	g.DELETE("/:id", controllers.RemoveTodo)
}
