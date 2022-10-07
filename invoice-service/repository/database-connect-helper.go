package repository

import (
	"database/sql"
	"fmt"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/internal/config"
	_ "github.com/lib/pq"
)

// Set up DB connection based on environment config
func SetupDbConnection(env *config.Config) (*sql.DB, error) {
	db, err := sql.Open(env.RepoDriver, BuildDbConnectionParam(env))
	if err != nil {
		return nil, err
	}
	return db, CheckDbConnection(db)
}

// Ping test the DB to ensure we have stable connection.
// Todo: Implement retry.
func CheckDbConnection(db *sql.DB) error {
	return db.Ping()
}

// Util function to build DB connection string based on config parameters.
func BuildDbConnectionParam(env *config.Config) string {
	if env == nil {
		fmt.Println("env Config is nil, returning empty string")
		return ""
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s connect_timeout=%d sslmode=%s",
		env.RepoHost, env.RepoPort, env.RepoUsername, env.RepoPassword, env.RepoName, env.RepoTimeout, env.SslMode)
}
