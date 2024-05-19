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
	app.POST("/servicers", handlers.AddServicer)
	app.PUT("/servicers", handlers.EditServicer)

	app.GET("/requests", handlers.ViewRequests)
	app.POST("/requests", handlers.AddRequest)
	app.PUT("/requests", handlers.EditRequest)

	app.GET("/clients", handlers.ViewClients)
	app.POST("/clients", handlers.AddClient)
	app.PUT("/clients", handlers.EditClient)

	app.Logger.Fatal(
		app.Start(os.Getenv("PORT")),
	)
}
