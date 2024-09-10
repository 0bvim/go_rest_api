package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go_rest_api/service/user"
	"github.com/gorilla/mux"
)

// APIServer represents an API server configuration.
//
// It contains essential components needed to run an API server, including
// the network address to listen on and a database connection for data operations.
type APIServer struct {
	// addr is the network address on which the API server listens.
	// It's typically in the format "host:port".
	addr string

	// db is a pointer to a SQL database connection.
	// It's used for storing and retrieving data accessed through the API.
	db *sql.DB
}

// NewAPIServer creates and returns a new instance of APIServer.
//
// It initializes an APIServer struct with the provided address and database connection.
//
// Parameters:
//
//	addr string: The network address on which the API server will listen.
//	            Typically in the format "host:port".
//	db *sql.DB: A pointer to a SQL database connection.
//	            This is likely used for storing and retrieving data accessed through the API.
//
// Returns:
//
//	*APIServer: A pointer to a newly created APIServer instance.
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// TODO: documentation
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
