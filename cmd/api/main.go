package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/xerdin442/api-practice/internal/env"
	repo "github.com/xerdin442/api-practice/internal/repository"
)

type application struct {
	port   int
	models repo.Registry
}

func main() {
	// Validate connection string
	db, err := sql.Open("mysql", env.GetStr("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatal("Invalid database connection string:", err)
	}

	// Connect to database
	if err := db.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}
	defer db.Close()

	// Configure app
	registry := repo.NewRegistry(db)
	app := &application{
		port:   env.GetInt("PORT"),
		models: *registry,
	}

	// Start the http server
	if err := app.serve(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
