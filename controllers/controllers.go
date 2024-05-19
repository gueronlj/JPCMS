package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/JPCMS/models"
	"github.com/labstack/echo"
)

func viewClients(data echo.Context) error {
	connectDB := "postgresql://gueronlj:4R3mijJzQYCc@ep-mute-term-70885178.us-east-2.aws.neon.tech/jps?sslmode=require"
	connection, err := sql.Open("postgres", connectDB)
	if err != nil {
		fmt.Println(err)
	}
	defer connection.Close()
	rows, err := connection.Query("SELECT * FROM clients;")
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
