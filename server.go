package main

import (
	"github.com/jorgeAM/echo-api/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", controllers.HiTodo)

	e.Logger.Fatal(e.Start(":8000"))
}
