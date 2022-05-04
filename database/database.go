package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		host, 5432, user, password, dbName)
	
	log.Printf("connecting to: %s", psqlInfo)
	
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("error when trying to connect to %s, %s", dbName, err.Error())
		return db, nil
	}
	
	return db, nil
}
