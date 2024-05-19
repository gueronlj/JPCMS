package handlers

import (
	"fmt"
	"net/http"

	"github.com/JPCMS/db"
	"github.com/JPCMS/models"
	"github.com/labstack/echo/v4"
)

func ViewClients(data echo.Context) error {
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

func CreateClient(client models.Client) (models.Client, error) {
	database := db.GetDB()
	query := `INSERT INTO clients (id, firstname, lastname, address) VALUES ($1, $2, $3, $4) RETURNING ID`
	err := database.QueryRow(query, client.ID, client.FirstName, client.LastName, client.Address).Scan(&client.ID)
	if err != nil {
		return client, err
	}
	return client, nil
}

func UpdateClient(client models.Client) (models.Client, error) {
	database := db.GetDB()
	query :=
		`UPDATE clients 
		SET FirstName = $2, LastName = $3, Address = $4
		WHERE id = $1
		RETURNING id`
	err := database.QueryRow(query, client.ID, client.FirstName, client.LastName, client.Address).Scan(&client.ID)
	if err != nil {
		return client, err
	}
	return client, nil
}

func EditClient(data echo.Context) error {
	client := models.Client{}
	data.Bind(&client)
	finalClient, err := UpdateClient(client)
	if err != nil {
		return data.JSON(http.StatusInternalServerError, err.Error())
	}
	return data.JSON(http.StatusCreated, finalClient)
}

func AddClient(data echo.Context) error {
	client := models.Client{}
	data.Bind(&client)
	newClient, err := CreateClient(client)
	if err != nil {
		return data.JSON(http.StatusInternalServerError, err.Error())
	}
	return data.JSON(http.StatusCreated, newClient)
}

func ViewServicers(data echo.Context) error {
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

func ViewRequests(data echo.Context) error {
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
