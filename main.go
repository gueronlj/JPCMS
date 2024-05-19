package main

import (
	"fmt"
	"os"

	"github.com/JPCMS/db"
	"github.com/JPCMS/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	//Start the http server
	fmt.Println("Running JPCMS backend on port 8080")
	db.InitDB()
	app := echo.New()

	app.GET("/servicers", handlers.ViewServicers)
	app.GET("/requests", handlers.ViewRequests)
	app.GET("/clients", handlers.ViewClients)
	app.POST("/clients", handlers.AddClient)
	app.PUT("/clients", handlers.EditClient)
	// prints error and exits program
	PORT := os.Getenv("PORT")
	app.Logger.Fatal(
		app.Start(PORT),
	)
}
