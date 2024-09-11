package main

// native import from github import normally are separated
import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/go_rest_api/cmd/api"
	"github.com/go_rest_api/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{})

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
