package database

import (
	"context"
	"database/sql"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

func (s *service) FindOrCreateUser(ctx context.Context, user *User) (string, error) {
	var userID string
	err := s.db.QueryRowContext(ctx, "SELECT id FROM users WHERE email = $1", user.Email).Scan(&userID)
	if err == sql.ErrNoRows {
		err = s.db.QueryRowContext(ctx, "INSERT INTO users (name, email, picture) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Email, user.Picture).Scan(&userID)
	}
	return userID, err
}

func (s *service) GetUserByID(ctx context.Context, userID string) (*User, error) {
	var user User
	err := s.db.QueryRowContext(ctx, "SELECT id, name, email, picture FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Name, &user.Email, &user.Picture)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}
