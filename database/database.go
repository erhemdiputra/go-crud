package database

import (
	"database/sql"
	"log"
)

var globalDB *sql.DB

func Init() error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-crud")
	if err != nil {
		return err
	}

	globalDB = db

	err = globalDB.Ping()
	if err != nil {
		return err
	}

	log.Println("Connect to MySQL successfully")
	return nil
}

func Get() *sql.DB {
	return globalDB
}
