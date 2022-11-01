package dbrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/inuoshios/keepinfo/internal/models"
)

// InsertUser inserts a user into the database
func (u *postgresDBRepo) InsertUser(user *models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	var newID int

	query := `
		INSERT INTO users (first_name, last_name, email, password, created_at, updated_at, deleted_at)
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	err := u.DB.QueryRowContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		time.Now(),
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, fmt.Errorf("error inserting into table %w", err)
	}

	return newID, nil
}

func (u *postgresDBRepo) GetUserbyEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	var user models.User

	query := `
	SELECT * FROM users WHERE email = $1`

	err := u.DB.QueryRowContext(ctx, query, email).Scan(&user)
	if errors.Is(err, sql.ErrNoRows) {
		return &user, fmt.Errorf("error: %w", err)
	}

	return &user, nil
}

// func (u *postgresDBRepo) LoginUser(username, password string) error {

// }
