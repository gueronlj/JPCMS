package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/JPCMS/db"
	"github.com/JPCMS/handlers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

var secretKey = []byte(os.Getenv("SECRET"))

func login(c echo.Context) error {
	// username := c.FormValue("username")
	// password := c.FormValue("password")
	// // Throws unauthorized error
	// if username != "jon" || password != "shhh!" {
	// 	return echo.ErrUnauthorized
	// }

	// Set custom claims
	claims := JwtCustomClaims{
		"jon",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenString,
	})
}

func main() {
	//Start the http server
	fmt.Println("Running JPCMS backend on port 8080")
	db.InitDB()

	// authConfig := echojwt.Config{
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(JwtCustomClaims)
	// 	},
	// 	SigningKey: []byte(secretKey),
	// }

	app := echo.New()
	//app.Use(echojwt.WithConfig(authConfig))
	app.POST("/login", login)
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
