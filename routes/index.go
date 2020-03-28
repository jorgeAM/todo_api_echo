package routes

import "github.com/labstack/echo/v4"

// InitRoutes init all routes
func InitRoutes(e *echo.Echo) {
	initAuthRoutes(e)
	initTodoRoutes(e)
}
