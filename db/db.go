package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gueronlj/JPCMS/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	connectionStr := os.Getenv("DB_STR")
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	connectionStr := os.Getenv("DB_STR")
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func CreateClient(client models.Client) (models.Client, error) {
	database := GetDB()
	query := `INSERT INTO clients (id, firstname, lastname, address) VALUES ($1, $2, $3, $4) RETURNING ID`
	err := database.QueryRow(query, client.ID, client.FirstName, client.LastName, client.Address).Scan(&client.ID)
	if err != nil {
		return client, err
	}
	return client, nil
}

func UpdateClient(client models.Client) (models.Client, error) {
	database := GetDB()
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

func CreateServicer(servicer models.Servicer) (models.Servicer, error) {
	database := GetDB()
	query := `INSERT INTO servicers (id, firstname, lastname, status) VALUES ($1, $2, $3, $4) RETURNING ID`
	err := database.QueryRow(query, servicer.ID, servicer.FirstName, servicer.LastName, servicer.Status).Scan(&servicer.ID)
	if err != nil {
		return servicer, err
	}
	return servicer, nil
}

func UpdateServicer(servicer models.Servicer) (models.Servicer, error) {
	database := GetDB()
	query :=
		`UPDATE servicers 
		SET FirstName = $2, LastName = $3, Status = $4
		WHERE id = $1
		RETURNING id`
	err := database.QueryRow(query, servicer.ID, servicer.FirstName, servicer.LastName, servicer.Status).Scan(&servicer.ID)
	if err != nil {
		return servicer, err
	}
	return servicer, nil
}

func CreateRequest(request models.Request) (models.Request, error) {
	database := GetDB()
	query := `INSERT INTO requests (id, ClientID, ServicerID, Address, InvoiceNumber, Description, completiondate, complettiontime) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING ID`
	err := database.QueryRow(query, request.ID, request.ClientID, request.ServicerID, request.Address, request.InvoiceNumber, request.Description, request.Date, request.Time).Scan(&request.ID)
	if err != nil {
		return request, err
	}
	return request, nil
}

func UpdateRequest(request models.Request) (models.Request, error) {
	database := GetDB()
	query :=
		`UPDATE requests 
		SET ClientID = $2, ServicerID = $3, Address = $4, InvoiceNumber = $5, Description = $6, completiondate = $7, complettiontime = $8
		WHERE id = $1
		RETURNING id`
	err := database.QueryRow(query, request.ID, request.ClientID, request.ServicerID, request.Address, request.InvoiceNumber, request.Description, request.Date, request.Time).Scan(&request.ID)
	if err != nil {
		return request, err
	}
	return request, nil
}
