package main

// native import from github import normally are separated
import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/go_rest_api/cmd/api"
	"github.com/go_rest_api/db"
	"github.com/go_rest_api/config"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
    User: config.Envs.DBUser,
    Passwd: config.Envs.DBPassword,
    Addr: config.Envs.DBAddress,
    DBName: config.Envs.DBName,
    AllowNativePassword: true,
    ParseTime: true,
  })
  if err != nil {
    log.Fatal(err)
  }

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
