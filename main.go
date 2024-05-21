package main

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/gueronlj/JPCMS/db"
	"github.com/gueronlj/JPCMS/handlers"
	"github.com/gueronlj/JPCMS/middleware"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	//Start the http server
	fmt.Println("Running JPCMS backend on port 8080")

	db.InitDB()

	app := echo.New()
	app.Renderer = &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}
	app.GET("/", handlers.Loginpage)
	app.POST("/signin", handlers.AttemptSignIn)

	app.GET("/servicers", handlers.ViewServicers)
	app.POST("/servicers", handlers.AddServicer)
	app.PUT("/servicers", handlers.EditServicer)

	app.GET("/requests", middleware.JWTAuth(handlers.ViewRequests))
	app.POST("/requests", handlers.AddRequest)
	app.PUT("/requests", handlers.EditRequest)

	app.GET("/clients", handlers.ViewClients)
	app.POST("/clients", handlers.AddClient)
	app.PUT("/clients", handlers.EditClient)

	app.Logger.Fatal(
		app.Start(os.Getenv("PORT")),
	)
}
