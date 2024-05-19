package db

import (
	"database/sql"
	"fmt"
	"os"

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
