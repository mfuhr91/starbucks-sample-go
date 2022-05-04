package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func DatabaseConnect() (*sql.DB, error) {
	var db *sql.DB
	var err error
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	if dbName == "" {
		dbName = "starbucks_db"
	}
	dataSource := fmt.Sprintf("%s:%s@tcp:(%s:3306)/%s", user, password, host, dbName)
	
	db, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("error when trying to connect to starbucks_db, %s", err.Error())
		return db, nil
	}
	
	return db, nil
}
