package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// This struct is expansible
type generalController struct{}

var GeneralController generalController

func (g generalController) Register(controller *echo.Group) {
	controller.GET("/", root)
}

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from the backend!")
}
