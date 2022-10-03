package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func SetupDatabase(driverName string, connString string) (*sql.DB, error) {
	// change "postgres" for whatever supported database you want to use
	db, err := sql.Open(driverName, connString)

	// Todo: duplicate code, better error handling.
	if err != nil {
		return nil, err
	}

	// ping the DB to ensure that it is connected
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func FetchRows(db *sql.DB) (*sql.Rows, error) {
	rows,err := db.Query("SELECT * from invoices where id = 1")
	fmt.Println("This is what we got !!!!")

	fmt.Println(rows)
if err != nil {
	log.Fatal(err)
}
return rows, err
}