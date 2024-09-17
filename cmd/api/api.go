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

// Run starts the HTTP server for the APIServer instance. (like a method)
//
// This method initializes a new Gorilla Mux router, sets up an API version 1
// subrouter with route handlers, and starts the server to listen on the
// specified address.
//
// The router handles API requests by delegating them to registered handlers.
// In this example, the `userHandler` is responsible for managing user-related
// routes under the `/api/v1` prefix.
//
// Log output will indicate the server's listening address.
//
// Returns:
// - `error`: Any error that occurs while starting or running the HTTP server.
//
// Example:
//
//	apiServer := &APIServer{addr: ":8080"}
//	err := apiServer.Run()
//	if err != nil {
//		log.Fatal(err)
//	}
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// dependence injection
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
