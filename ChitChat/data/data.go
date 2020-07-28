package data

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
