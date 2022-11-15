package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/ixxiv/keepinfo/internal/models"
	"github.com/ixxiv/keepinfo/internal/utils"
)

func (u *postgres) InsertContact(contact *models.Contact) (string, error) {
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
		return "", fmt.Errorf("error inserting into contacts table %w", err)
	}

	return string(newId), nil

}

func (u *postgres) GetContacts(args models.GetAllUsers) ([]models.Contact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	var contact []models.Contact

	query := `
	SELECT id, user_id, first_name, last_name, email, phone, label, address, created_at, updated_at
	FROM contacts 
	WHERE user_id = $1 
	ORDER BY first_name`

	rows, err := u.DB.QueryContext(ctx, query, args.UserID)
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

func (u *postgres) GetContact(id string) (models.Contact, error) {
	var contact models.Contact
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	query := `
	SELECT id, user_id, first_name, last_name, email, phone, label, address, created_at, updated_at
	FROM contacts WHERE id = $1 LIMIT 1`

	rows := u.DB.QueryRowContext(ctx, query, id)
	err := rows.Scan(
		&contact.ID,
		&contact.UserID,
		&contact.FirstName,
		&contact.LastName,
		&contact.Email,
		&contact.Phone,
		&contact.Label,
		&contact.Address,
		&contact.CreatedAt,
		&contact.UpdatedAt,
	)

	return contact, err
}

func (u *postgres) UpdateContact(contact *models.Contact) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	query := `
	UPDATE contacts
	SET first_name = $1, last_name = $2, email = $3, phone = $4, label = $5, address = $6,
	updated_at = $7, user_id = $8
	WHERE id = $9
	RETURNING id
	`
	rows := u.DB.QueryRowContext(ctx, query,
		contact.FirstName,
		contact.LastName,
		contact.Email,
		contact.Phone,
		contact.Label,
		contact.Address,
		contact.UpdatedAt,
		contact.UserID,
		contact.ID,
	)
	err := rows.Scan(&contact.ID)

	return err
}

func (u *postgres) DeleteContact(id, userid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	query := `
	DELETE FROM contacts
	WHERE id = $1 AND user_id = $2
	`

	rows, err := u.DB.ExecContext(ctx, query, id, userid)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return utils.ErrContactSqlNoRows
	}

	return nil
}
