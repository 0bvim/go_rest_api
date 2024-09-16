package user

import (
	"net/http"

	"github.com/go_rest_api/types"
	"github.com/go_rest_api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

// NewHandler create and return a Handler instance
func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registers the routes for handling user login and registration.
//
// This method sets up two POST routes for the given router. It connects the "/login" route to the
// handleLogin handler and the "/register" route to the handleRegister handler.
//
// Parameters:
//
//	router (*mux.Router): The router from the Gorilla Mux package to which the routes are registered.
//
// Example Usage:
//
//	router := mux.NewRouter()
//	handler := &Handler{}
//	handler.RegisterRoutes(router)
//
// The routes will be accessible at the following endpoints:
//
//	POST /login
//	POST /register
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// check if the user exists
	// if it doesnt we create the new user
}
