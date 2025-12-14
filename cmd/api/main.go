package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type application struct {
	port string
	// models
}

func main() {
	_, err := sql.Open("mysql", os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		panic(err)
	}

	app := &application{
		port: os.Getenv("PORT"),
	}

	if err := app.serve(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
