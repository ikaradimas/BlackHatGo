package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	dsn := "postgres://postgres:password@localhost:5432/store?sslmode=disable"
	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	defer db.Close()

	var (
		ccnum, date, cvv, exp string
		amount                float32
	)

	rows, err := db.Query("SELECT ccnum, date, cvv, exp, amount FROM transactions")
	if err != nil {
		log.Panicln(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&ccnum, &date, &amount, &cvv, &exp)
		if err != nil {
			log.Panicln(err)
		}

		fmt.Println(ccnum, date, amount, cvv, exp)
		if rows.Err() != nil {
			log.Panicln(err)
		}
	}
}
