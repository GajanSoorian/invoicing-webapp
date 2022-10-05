package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/internal/config"
	_ "github.com/lib/pq"
)

func SetupDatabase(env *config.Config) (*sql.DB, error) {
	db, err := sql.Open(env.RepoDriver, BuildDbConnectionParam(env))
	if err != nil {
		return nil, err
	}
	return db, checkDbConnection(db)
}

// Ping test the DB to ensure we have stable connection.
// Todo: Implement retry.
func checkDbConnection(db *sql.DB) error {
	return db.Ping()
}

//Todo 
func FetchRows(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT * from invoices where id = 1")
	fmt.Println("This is what we got !!!!")
	fmt.Println(rows)
	if err != nil {
		log.Fatal(err)
	}
	return rows, err
}

func BuildDbConnectionParam(env *config.Config) string {
	if env == nil {
		fmt.Println("env Config is nil, returning empty string")
		return ""
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s connect_timeout=%d sslmode=%s",
		env.RepoHost, env.RepoPort, env.RepoUsername, env.RepoPassword, env.RepoName, env.RepoTimeout, env.SslStatus)
}
