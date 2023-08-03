package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/jackjohn7/go_svelte_spa_template/controllers"
)

func main() {
	app := echo.New()
	app.Pre(middleware.AddTrailingSlash())

	// add SPA (SvelteKit app)
	app.Static("/", "client/build")

	// create handler group for backend functions
	apiGroup := app.Group("/api")
	// create controllers with apiGroup
	controllers.CreateGeneralController(apiGroup)

	go func() {
		if err := app.Start(":5173"); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal(err)
		}
	}()

	// wait for exit signal (allows for graceful shutdowns)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}
}
