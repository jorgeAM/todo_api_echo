package routes

import (
	"github.com/jorgeAM/echo-api/controllers"
	"github.com/labstack/echo/v4"
)

func initAuthRoutes(e *echo.Echo) {
	g := e.Group("auth")
	g.POST("/sign-up", controllers.SignUp)
	g.POST("/login", controllers.Login)
}
