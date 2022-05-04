package database

import (
	"database/sql"
	"log"
)

func DatabaseConnect() (*sql.DB, error) {
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", "admin:2744mfuhR@/starbucks_db")
	if err != nil {
		log.Fatalf("error when trying to connect to starbucks_db, %s", err.Error())
		return db, nil
	}
	
	return db, nil
}
