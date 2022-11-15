package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/ixxiv/keepinfo/internal/models"
)

// InsertUser inserts a user into the database
func (u *postgres) InsertUser(user *models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	query := `
		INSERT INTO users (username, first_name, last_name, email, password, created_at, updated_at, deleted_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING username`

	err := u.DB.QueryRowContext(ctx, query,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
	).Scan(&user.Username)

	if err != nil {
		return "", fmt.Errorf("error inserting into user table %w", err)
	}

	return user.Username, nil
}

func (u *postgres) GetUser(email string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	query := `
	SELECT username, first_name, last_name, email, password, created_at, updated_at
	FROM users WHERE email = $1`

	rows := u.DB.QueryRowContext(ctx, query, email)

	err := rows.Scan(
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}

func (u *postgres) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	query := `
	SELECT username, first_name, last_name, email, password, created_at, updated_at
	FROM users WHERE username = $1`

	rows := u.DB.QueryRowContext(ctx, query, username)

	err := rows.Scan(
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}
