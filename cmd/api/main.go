package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load environment variables")
	}

	_, err := sql.Open("mysql", os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	if err := r.Run(os.Getenv("PORT")); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
