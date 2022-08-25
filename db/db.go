package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func DBConnect() *sql.DB {

	var err error
	connStr := "user= password= dbname= sslmode=disable"
	pgUrl, _ := pq.ParseURL("connStr")

	db, err := sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}
	return db
}
