package main

import (
	"github.com/jackjohn7/go_svelte_spa_template/controllers"
	"github.com/jackjohn7/go_svelte_spa_template/server"
)

// add any SPA clients here
var clients []*server.Client = []*server.Client{
	{
		Name:      "client",
		Prefix:    "/",
		OutputDir: "client/build",
		BuildCmd:  []string{"npm", "run", "build"},
	},
}

func main() {
	// watch states that when changes are detected in the clients, they're rebuilt
	watch := true
	server := server.CreateServer(5173, clients)

	// register the general controller (any struct implementing the Controller interface)
	server.RegisterController("/api", controllers.GeneralController)

	server.Start(watch)
}
