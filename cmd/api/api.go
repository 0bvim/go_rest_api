package api

// import necessary packages to api and database connection
import "database/sql"

// struct of server api
type APIServer struct {
	addr string
	db   *sql.DB
}

// function that return a pointer to an APIServer
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
