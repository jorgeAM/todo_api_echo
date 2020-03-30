package routes

import (
	"os"

	"github.com/jorgeAM/echo-api/controllers"
	"github.com/jorgeAM/echo-api/middlewares"
	"github.com/labstack/echo/v4"
)

var key = os.Getenv("JWT_KEY")

func initTodoRoutes(e *echo.Echo) {
	g := e.Group("todos")
	g.GET("", controllers.GetTodos, middlewares.AuthMiddleware())
	g.POST("", controllers.NewTodo, middlewares.AuthMiddleware())
	g.GET("/:id", controllers.GetTodo)
	g.PUT("/:id", controllers.UpdateTodo)
	g.DELETE("/:id", controllers.RemoveTodo)
}
