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

func (u *postgresDBRepo) GetContacts() ([]models.Contact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	var contact []models.Contact

	query := `
	SELECT id, user_id, first_name, last_name, email, phone, label, address, created_at, updated_at
	FROM contacts ORDER BY created_at`

	rows, err := u.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		rows.Close()
	}()

	for rows.Next() {
		var c models.Contact
		err := rows.Scan(
			&c.ID,
			&c.UserID,
			&c.FirstName,
			&c.LastName,
			&c.Email,
			&c.Phone,
			&c.Label,
			&c.Address,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		contact = append(contact, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return contact, nil
}
