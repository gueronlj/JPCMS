package main

import (
	"fmt"

	"github.com/JPCMS/handlers"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	//Start the http server
	fmt.Println("Running JPCMS backend on port 8080")
	app := echo.New()

	app.GET("/servicers", handlers.ServicerHandler{}.ViewServicers)
	app.GET("/requests", handlers.RequestHandler{}.ViewRequests)
	app.GET("/clients", handlers.ClientHandler{}.ViewClients)
	// prints error and exits program
	app.Logger.Fatal(
		app.Start(":8080"),
	)
}
