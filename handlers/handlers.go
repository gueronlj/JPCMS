package handlers

import (
	"fmt"
	"net/http"

	"github.com/JPCMS/db"
	"github.com/JPCMS/models"
	"github.com/labstack/echo/v4"
)

type ClientHandler struct{}
type ServicerHandler struct{}
type RequestHandler struct{}

func (h ClientHandler) ViewClients(data echo.Context) error {
	database := db.GetDB()
	rows, err := database.Query("SELECT * FROM clients;")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var clients []models.Client
	for rows.Next() {
		var cli models.Client
		err := rows.Scan(&cli.ID, &cli.FirstName, &cli.LastName, &cli.Address)
		if err != nil {
			panic(err)
		}
		clients = append(clients, cli)
	}
	return data.JSON(http.StatusOK, clients)
}

func (h ServicerHandler) ViewServicers(data echo.Context) error {
	database := db.GetDB()
	rows, err := database.Query("SELECT * FROM servicers;")
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

func (h RequestHandler) ViewRequests(data echo.Context) error {
	database := db.GetDB()
	rows, err := database.Query("SELECT * FROM requests;")
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
