package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go_rest_api/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		// creating payload to test
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "abc",
			Email:     "",
			Password:  "123",
		}

		// enconding payload
		marshalled, _ := json.Marshal(payload)
		// mock the request with encoded payload by marshal
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		requestRecorder := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(requestRecorder, req)
		if requestRecorder.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d",
				http.StatusBadRequest, requestRecorder.Code)
		}
	})
}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
