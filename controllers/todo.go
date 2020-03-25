package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HiTodo says hi to you
func HiTodo(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hi!!",
	})
}
