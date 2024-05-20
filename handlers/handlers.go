package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gueronlj/JPCMS/auth"
	"github.com/gueronlj/JPCMS/db"
	"github.com/gueronlj/JPCMS/models"
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

func EditClient(data echo.Context) error {
	client := models.Client{}
	data.Bind(&client)
	finalClient, err := db.UpdateClient(client)
	if err != nil {
		return data.JSON(http.StatusInternalServerError, err.Error())
	}
	return data.JSON(http.StatusCreated, finalClient)
}

func AddClient(data echo.Context) error {
	client := models.Client{}
	data.Bind(&client)
	newClient, err := db.CreateClient(client)
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

func EditServicer(data echo.Context) error {
	servicer := models.Servicer{}
	data.Bind(&servicer)
	finalServicer, err := db.UpdateServicer(servicer)
	if err != nil {
		return data.JSON(http.StatusInternalServerError, err.Error())
	}
	return data.JSON(http.StatusCreated, finalServicer)
}

func AddServicer(data echo.Context) error {
	servicer := models.Servicer{}
	data.Bind(&servicer)
	newServicer, err := db.CreateServicer(servicer)
	if err != nil {
		return data.JSON(http.StatusInternalServerError, err.Error())
	}
	return data.JSON(http.StatusCreated, newServicer)
}

func ViewRequests(data echo.Context) error {
	if auth.CheckJWT(data.Request()) == "success" {
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
	return data.JSON(http.StatusUnauthorized, auth.CheckJWT(data.Request()))
}

func EditRequest(data echo.Context) error {
	request := models.Request{}
	data.Bind(&request)
	finalRequest, err := db.UpdateRequest(request)
	if err != nil {
		return data.JSON(http.StatusInternalServerError, err.Error())
	}
	return data.JSON(http.StatusCreated, finalRequest)
}

func AddRequest(data echo.Context) error {
	request := models.Request{}
	data.Bind(&request)
	newRequest, err := db.CreateRequest(request)
	if err != nil {
		return data.JSON(http.StatusInternalServerError, err.Error())
	}
	return data.JSON(http.StatusCreated, newRequest)
}

func Login(c echo.Context) error {
	tokenString, err := auth.CreateToken("Jon")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tokenString)
	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenString,
	})
}

// func Login(c echo.Context) error {
// 	// Set custom claims
// 	claims := JwtCustomClaims{
// 		"jon",
// 		true,
// 		jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte(secretKey))
// 	if err != nil {
// 		return err
// 	}
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"token": tokenString,
// 	})
// }
