package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// SignIn creates a new user
func SignIn(c echo.Context) error {
	return c.String(http.StatusOK, "Sign in")
}

// Login returns a jwt
func Login(c echo.Context) error {
	return c.String(http.StatusOK, "Login")
}
