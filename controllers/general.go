package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateGeneralController(api *echo.Group) {
	api.GET("/", root)
}

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from the backend!")
}
