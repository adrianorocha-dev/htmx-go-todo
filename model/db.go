package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Setup() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=myuser dbname=mydatabase password=mypassword sslmode=disable")

	if err != nil {
		fmt.Println("Could not connect to the db")
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Could not ping the db")
		panic(err)
	}
}
