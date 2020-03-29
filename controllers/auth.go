package controllers

import (
	"net/http"

	connection "github.com/jorgeAM/echo-api/db/conn"
	"github.com/jorgeAM/echo-api/models"
	"github.com/jorgeAM/echo-api/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// SignUp creates a new user
func SignUp(c echo.Context) error {
	db := connection.Conn()
	defer db.Close()
	u := new(models.User)

	if err := c.Bind(&u); err != nil {
		return err
	}

	if u.Password != u.ConfirmPassword {
		return echo.NewHTTPError(http.StatusBadRequest, "passwords don't match")
	}

	db.Where("email = ?", u.Email).First(&u)

	if u.ID > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "email is already taken")
	}

	hash, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	u.Password = hash

	if err := db.Create(&u).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "we can't create your account")
	}

	return c.JSON(http.StatusCreated, &u)
}

// Login returns a jwt
func Login(c echo.Context) error {
	db := connection.Conn()
	defer db.Close()
	u := new(models.User)

	if err := c.Bind(&u); err != nil {
		return err
	}

	pwd := u.Password

	db.Where("email = ?", u.Email).First(&u)

	if u.ID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "account does not exist")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect password")
	}

	return c.JSON(http.StatusOK, &u)
}
