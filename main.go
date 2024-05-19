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
	app := echo.New()

	app.GET("/servicers", handlers.ServicerHandler{}.ViewServicers)
	app.GET("/requests", handlers.RequestHandler{}.ViewRequests)
	app.GET("/clients", handlers.ClientHandler{}.ViewClients)
	// prints error and exits program
	db.InitDB()
	PORT := os.Getenv("PORT")
	app.Logger.Fatal(
		app.Start(PORT),
	)
}
