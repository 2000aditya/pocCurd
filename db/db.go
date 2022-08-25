package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func DBConnect() *sql.DB {

	var err error
	//connStr := "user=wuwojvci password=nuk1BuktGzd_k3jf5uX2vddsJ0alYHfZ dbname=wuwojvci sslmode=disable"
	pgUrl, _ := pq.ParseURL("postgres://wuwojvci:nuk1BuktGzd_k3jf5uX2vddsJ0alYHfZ@satao.db.elephantsql.com/wuwojvci?sslmode=disable")

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
