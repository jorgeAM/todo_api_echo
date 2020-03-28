package routes

import (
	"github.com/jorgeAM/echo-api/controllers"
	"github.com/labstack/echo/v4"
)

func initAuthRoutes(e *echo.Echo) {
	g := e.Group("auth")
	g.POST("/sign-in", controllers.SignIn)
	g.POST("/login", controllers.Login)
}
