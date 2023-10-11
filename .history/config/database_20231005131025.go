package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/go_product?parseTi")
	if err != nil {
		panic(err)
	}

	log.Printf("Database connected")
	DB = db
}