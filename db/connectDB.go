package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var db *sql.DB

func init() {
	var dbError error
	db, dbError = sql.Open("mysql", os.Getenv("DSN"))
	if dbError != nil {
		log.Fatalf("failed to connect: %v", dbError)
	}
}

func GetDB() *sql.DB {
	return db
}
