package user

import (
	"database/sql"
	"fmt"
	"github.com/emaanmohamed/shop/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowInToUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("User not found")
	}
	return u, nil
}

func scanRowInToUser(rows *sql.Rows) (*types.User, error) {
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
	return user, err
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (first_name, last_name, email, password, created_at) VALUES (?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
