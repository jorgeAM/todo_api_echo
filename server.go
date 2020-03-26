package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/echo-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
