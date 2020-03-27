package controllers

import (
	"net/http"

	connection "github.com/jorgeAM/echo-api/db/conn"
	"github.com/jorgeAM/echo-api/models"
	"github.com/labstack/echo/v4"
)

// GetTodos gets all todos
func GetTodos(c echo.Context) error {
	db := connection.Conn()
	defer db.Close()
	todos := []models.Todo{}
	db.Find(&todos)
	return c.JSON(http.StatusOK, &todos)
}

// GetTodo returns a todo by id
func GetTodo(c echo.Context) error {
	db := connection.Conn()
	defer db.Close()
	t := models.Todo{}
	id := c.Param("id")
	db.First(&t, id)

	if t.ID <= 0 {
		msg := "todo with id " + id + " does not exist"
		return echo.NewHTTPError(http.StatusNotFound, msg)
	}

	return c.JSON(http.StatusOK, &t)
}

// NewTodo creates new todo
func NewTodo(c echo.Context) error {
	db := connection.Conn()
	defer db.Close()
	t := models.Todo{}
	err := c.Bind(&t)

	if err != nil {
		return err
	}

	err = db.Create(&t).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &t)
}

// UpdateTodo updates a todo
func UpdateTodo(c echo.Context) error {
	db := connection.Conn()
	defer db.Close()
	t := models.Todo{}
	id := c.Param("id")
	db.First(&t, id)

	if t.ID <= 0 {
		msg := "user with id " + id + " does not exist"
		return echo.NewHTTPError(http.StatusNotFound, msg)
	}

	err := c.Bind(&t)

	if err != nil {
		return err
	}

	db.Save(&t)
	return c.JSON(http.StatusOK, &t)
}

// RemoveTodo deletes a todo
func RemoveTodo(c echo.Context) error {
	db := connection.Conn()
	defer db.Close()
	t := models.Todo{}
	id := c.Param("id")
	db.First(&t, id)

	if t.ID <= 0 {
		msg := "user with id " + id + " does not exist"
		return echo.NewHTTPError(http.StatusNotFound, msg)
	}

	err := db.Delete(&t).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}
