package controllers

import (
	"net/http"

	connection "github.com/jorgeAM/echo-api/db/conn"
	"github.com/jorgeAM/echo-api/models"
	"github.com/labstack/echo/v4"
)

// HiTodo says hi to you
func HiTodo(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hi!!",
	})
}

// GetTodo returns a todo by id
func GetTodo(c echo.Context) error {
	db := connection.Conn()
	defer db.Close()
	u := models.Todo{}
	id := c.Param("id")
	db.First(&u, id)

	if u.ID <= 0 {
		msg := "user with id " + id + " does not exist"
		return echo.NewHTTPError(http.StatusNotFound, msg)
	}

	return c.JSON(http.StatusOK, &u)
}

// NewTodo creates new todo
func NewTodo(c echo.Context) error {
	db := connection.Conn()
	defer db.Close()
	u := models.Todo{}
	err := c.Bind(&u)

	if err != nil {
		return err
	}

	err = db.Create(&u).Error

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
