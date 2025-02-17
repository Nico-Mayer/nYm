package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/nico-mayer/nym/config"
)

var DB *sql.DB

func Connect() error {
	var err error
	DB, err = sql.Open("postgres", config.DATABASE_URL)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}
