package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Represents a client application in the form of an SPA
type Client struct {
	sync.RWMutex
	// directory of client application itself
	Name string
	// prefix for routes
	Prefix string
	// location of build files
	OutputDir string
	// command to build project
	// e.g {"npm", "run", "build"}
	BuildCmd []string
}

func (c *Client) Build() (output string, err error) {
	c.RLock()
	cmd := exec.Command(c.BuildCmd[0], c.BuildCmd[1:]...)
	cmd.Dir = c.Name
	var outputRaw []byte
	outputRaw, err = cmd.Output()
	output = string(outputRaw)
	return
}

// Represents the application in its entirety
type Server struct {
	// Contains a pointer to an Echo web server
	App *echo.Echo
	// All of your SPA clients. Can be empty
	Clients []*Client
	// Port where the entire app is exposed
	Port string
}

type Controller interface {
	Register(e *echo.Group)
}

func CreateServer(port int, clients []*Client) *Server {
	app := echo.New()
	app.Pre(middleware.AddTrailingSlash())

	// serve static build files
	for _, c := range clients {
		output, err := c.Build()
		if err != nil {
			fmt.Println("Build command failed. Error output below:")
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(output)
		app.Static(c.Prefix, c.OutputDir)
	}

	return &Server{
		App:     app,
		Clients: clients,
		Port:    fmt.Sprintf(":%d", port),
	}
}

func (s *Server) ReloadClientsOnChanges() {
}

func (s *Server) RegisterController(prefix string, c Controller) {
	c.Register(s.App.Group(prefix))
}

func (s *Server) Start(watch bool) {
	// if user elects to watch for changes, reload clients on modifications
	if watch {
		go s.ReloadClientsOnChanges()
	}

	go func() {
		if err := s.App.Start(s.Port); err != nil && err != http.ErrServerClosed {
			s.App.Logger.Fatal(err)
		}
	}()
	// wait for exit signal (allows for graceful shutdowns)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.App.Shutdown(ctx); err != nil {
		s.App.Logger.Fatal(err)
	}
}
