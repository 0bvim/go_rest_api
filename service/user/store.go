package user

// this is a storage layer
// this is like repositories pattern

import (
	"database/sql"
	"fmt"

	"github.com/go_rest_api/types"
)

// data base used to make queries and so forth
type Store struct {
	db *sql.DB
}

// function to get db
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// method to get user using email string
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not foun")
	}

	return user, nil
}

// function that search in rows of db for the user information
func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
