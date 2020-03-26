package controllers

import (
	"net/http"

	"github.com/jorgeAM/echo-api/db"
	"github.com/jorgeAM/echo-api/models"
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

// GetTodo returns a todo by id
func GetTodo(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"message1": id,
	})
}

// NewTodo creates new todo
func NewTodo(c echo.Context) error {
	u := models.Todo{}
	err := c.Bind(&u)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &u)
}

// UpdateTodo updates a todo
func UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"message2": id,
	})
}

// RemoveTodo deletes a todo
func RemoveTodo(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"message3": id,
	})
}
