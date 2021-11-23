package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DbConecction() (db *sql.DB) {

	dbRetries := 30
	db, err := sql.Open("mysql", "root:secret@tcp(database)/gocker")
	if err != nil {
		log.Println(err.Error())
		for i := 0; i < dbRetries; i++ {
			time.Sleep(time.Second * time.Duration(i+1))
			db, err := sql.Open("mysql", "root:secret@tcp(database)/gocker")
			if err != nil {
				log.Printf("Error: %s retrying: %v", err.Error(), i)
				continue
			}
			return db
		}
	}
	return db
}
