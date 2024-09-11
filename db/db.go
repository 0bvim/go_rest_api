package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func NewPostgresStorage(cfg postgres.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
