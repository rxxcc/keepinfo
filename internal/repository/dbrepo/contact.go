package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/inuoshios/keepinfo/internal/models"
)

func (u *postgresDBRepo) InsertContact(contact *models.Contact) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	var newId []byte

	query := `
	INSERT INTO contacts (user_id, first_name, last_name, email, phone, label, address, created_at, updated_at, deleted_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id`

	err := u.DB.QueryRowContext(ctx, query,
		contact.UserID,
		contact.FirstName,
		contact.LastName,
		contact.Email,
		contact.Phone,
		contact.Label,
		contact.Address,
		contact.CreatedAt,
		contact.UpdatedAt,
		contact.DeletedAt,
	).Scan(&newId)

	if err != nil {
		return "", fmt.Errorf("error inserting into table %w", err)
	}

	return string(newId), nil

}
