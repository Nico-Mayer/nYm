package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/nico-mayer/nym/config"
)

var DB *sql.DB

func Connect() error {
	connStr := "user=" + config.PGUSER + " password=" + config.PGPASSWORD + " dbname=" + config.PGDATABASE + " host=" + config.PGHOST + " sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}
