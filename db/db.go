package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

// TODO: implement with postgres later
// type Config struct {
// 	Ussername  string
// 	Password   string
// 	Host       string
// 	Port       int
// 	Database   string
// 	Parameters map[string]string
// }

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
