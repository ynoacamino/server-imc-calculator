package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var dbError error
	db, dbError = sql.Open("mysql", "jscoycrtxw6be9giz82t:pscale_pw_7caNNowZePYqcR60dz72pChxsBZysYJLzaRGi49AeJ6@tcp(aws.connect.psdb.cloud)/todo-go?tls=true&interpolateParams=true")
	if dbError != nil {
		log.Fatalf("failed to connect: %v", dbError)
	}
}

func GetDB() *sql.DB {
	return db
}
