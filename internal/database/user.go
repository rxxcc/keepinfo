package database

import (
	"context"
	"time"

	"github.com/inuoshios/keepinfo/internal/models"
)

func (db *DB) InsertUser(user models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	var id int

	query := `
		INSERT INTO users (first_name, last_name, email, password, created_at, updated_at, deleted_at)
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	args := []interface{}{
		user.FirstName, user.LastName, user.Email, user.Password, time.Now(), time.Now(), time.Now(),
	}

	err := db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}
