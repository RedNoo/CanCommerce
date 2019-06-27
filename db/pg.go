package dbUtil

import (
	"database/sql"
	"log"
)

var Db_ *(sql.DB)

func Connect() {
	var err error
	//db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=cancommerce password=1234 sslmode=disable")

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
