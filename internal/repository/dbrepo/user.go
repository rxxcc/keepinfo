package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/inuoshios/keepinfo/internal/models"
)

// InsertUser inserts a user into the database
func (u *postgresDBRepo) InsertUser(user *models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	var newID []byte

	query := `
		INSERT INTO users (first_name, last_name, email, password, created_at, updated_at, deleted_at)
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	err := u.DB.QueryRowContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&newID)

	if err != nil {
		return "", fmt.Errorf("error inserting into table %w", err)
	}

	return string(newID), nil
}

func (u *postgresDBRepo) GetUser(email string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	query := `
	SELECT id, first_name, last_name, email, password, created_at, updated_at
	FROM users WHERE email = $1`

	rows := u.DB.QueryRowContext(ctx, query, email)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}
