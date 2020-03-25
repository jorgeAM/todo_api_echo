package main

import (
	"github.com/jorgeAM/echo-api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
