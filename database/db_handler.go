package database

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/hashtag_generator?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
