package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/inuoshios/keepinfo/internal/models"
)

type DB struct {
	DB *sql.DB
}

func (db *DB) InsertUser(user *models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int

	query := `
		INSERT INTO users (first_name, last_name, email, password, )
		VALUES($1, $2, $3, $4)
		RETURNING id`

	args := []interface{}{
		user.FirstName, user.LastName, user.Email, user.Password, time.Now(), time.Now(), time.Now(),
	}

	if err := db.DB.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (db *DB) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, first_name, last_name, email, password, created_at, updated_at, deleted_at
			from users where id = $1`

	row := db.DB.QueryRowContext(ctx, query, id)

	var data models.User
	err := row.Scan(
		&data.ID,
		&data.FirstName,
		&data.LastName,
		&data.Email,
		&data.Password,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	if err != nil {
		return &data, err
	}

	return &data, nil
}
