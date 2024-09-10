package main

// native import from github import normally are separated
import (
	"log"

	"github.com/go_rest_api/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
