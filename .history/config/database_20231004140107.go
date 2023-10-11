package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var 

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/go_products")
	if err != nil {
		panic(err)
	}
}