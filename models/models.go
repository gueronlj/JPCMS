package models

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
	Date          string
	Time          string
}
