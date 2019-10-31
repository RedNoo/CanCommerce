package dbUtil

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db_ *(sql.DB)

func Connect() {
	var err error

	connStr := "postgres://postgres:1234@localhost/cancommerce?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	Db_ = db
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

}
