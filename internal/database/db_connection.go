package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConecction() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:rugbychampagne2021@tcp(127.0.0.1:3306)/cts")
	if err != nil {

	}
	return db
}
