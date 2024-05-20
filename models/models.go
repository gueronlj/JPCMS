package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Servicer struct {
	ID        string
	FirstName string
	LastName  string
	Status    string
}

type Client struct {
	ID        string
	FirstName string
	LastName  string
	Address   string
}

type Request struct {
	ID            string
	ClientID      string
	ServicerID    string
	Address       string
	InvoiceNumber string
	Description   string
	Date          time.Time
	Time          time.Time
}

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}
