package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/xerdin442/api-practice/internal/cache"
	"github.com/xerdin442/api-practice/internal/env"
	repo "github.com/xerdin442/api-practice/internal/repository"
	"github.com/xerdin442/api-practice/internal/service"
)

type application struct {
	port     int
	services *service.Manager
	cache    *cache.Redis
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

	// Initialize repositories and services
	registry := repo.NewRegistry(db)
	services := service.NewManager(registry)

	// Initialize cache
	cache := cache.NewRedis()

	app := &application{
		port:     env.GetInt("PORT"),
		services: services,
		cache:    cache,
	}

	// Start the http server
	if err := app.serve(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
