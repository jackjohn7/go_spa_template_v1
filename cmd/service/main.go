package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})

	app.Static("/", "client/build")

	app.Start(":5173")
}
