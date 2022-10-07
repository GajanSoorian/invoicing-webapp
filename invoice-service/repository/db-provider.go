package repository

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/internal/config"
)

var Lock = &sync.Mutex{}

// Singleton DB handle- sql.DB is part of Go standard library.
var db *sql.DB

// Return db handle if it exists, or create a new handle and return it.
func GetDbHandle(env *config.Config) *sql.DB {
	var err error
	if db == nil {
		Lock.Lock()
		defer Lock.Unlock()
		if db == nil {
			db, _ = sql.Open(env.RepoDriver, BuildDbConnectionParam(env))
			if err != nil {
				fmt.Println("Uh oh error with db connection open")
			} else {
				fmt.Println("db connection open, sucess!")
			}
		}
	}
	return db
}

func Init(env *config.Config) {
	_ = GetDbHandle(env)
}
