package controllers

import (
	"net/http"

	"github.com/jorgeAM/echo-api/db"
	"github.com/labstack/echo/v4"
)

// HiTodo says hi to you
func HiTodo(c echo.Context) error {
	db := db.Conn()
	db.Close()
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hi!!",
	})
}
