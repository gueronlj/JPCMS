package main

import (
	"fmt"
	"os"

	"github.com/JPCMS/db"
	"github.com/JPCMS/handlers"
	"github.com/JPCMS/middleware/auth"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func main() {
	//Start the http server
	fmt.Println("Running JPCMS backend on port 8080")
	db.InitDB()

	authConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	app := echo.New()
	app.Use(echojwt.WithConfig(authConfig))

	app.GET("/servicers", handlers.ViewServicers)
	app.POST("/servicers", handlers.AddServicer)
	app.PUT("/servicers", handlers.EditServicer)

	app.GET("/requests", handlers.ViewRequests)
	app.POST("/requests", handlers.AddRequest)
	app.PUT("/requests", handlers.EditRequest)

	app.GET("/clients", handlers.ViewClients)
	app.POST("/clients", handlers.AddClient)
	app.PUT("/clients", handlers.EditClient)

	app.POST("/login", auth.GenerateToken())

	app.Logger.Fatal(
		app.Start(os.Getenv("PORT")),
	)
}
