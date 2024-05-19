package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/JPCMS/models"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func viewServicers(data echo.Context) error {
	connectDB := "postgresql://gueronlj:4R3mijJzQYCc@ep-mute-term-70885178.us-east-2.aws.neon.tech/jps?sslmode=require"
	connection, err := sql.Open("postgres", connectDB)
	if err != nil {
		fmt.Println(err)
	}
	defer connection.Close()

	rows, err := connection.Query("SELECT * FROM servicers;")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var servicers []models.Servicer
	for rows.Next() {
		var servicer models.Servicer
		err := rows.Scan(&servicer.ID, &servicer.FirstName, &servicer.LastName, &servicer.Status)
		if err != nil {
			panic(err)
		}
		servicers = append(servicers, servicer)
	}
	return data.JSON(http.StatusOK, servicers)
}

func viewRequests(data echo.Context) error {
	connectDB := "postgresql://gueronlj:4R3mijJzQYCc@ep-mute-term-70885178.us-east-2.aws.neon.tech/jps?sslmode=require"
	connection, err := sql.Open("postgres", connectDB)
	if err != nil {
		fmt.Println(err)
	}
	defer connection.Close()

	rows, err := connection.Query("SELECT * FROM requests;")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var requests []models.Request
	for rows.Next() {
		var req models.Request
		err := rows.Scan(&req.ID, &req.ClientID, &req.ServicerID, &req.Address, &req.InvoiceNumber, &req.Description, &req.Date, &req.Time)
		if err != nil {
			panic(err)
		}
		requests = append(requests, req)
	}
	return data.JSON(http.StatusOK, requests)
}

func main() {
	//Start the http server
	fmt.Println("Running JPCMS backend on port 8080")
	app := echo.New()

	app.GET("/servicers", viewServicers)
	app.GET("/requests", viewRequests)
	app.GET("/clients", controllers.viewClients)

	// prints error and exits program
	app.Logger.Fatal(
		app.Start(":8080"),
	)
}
