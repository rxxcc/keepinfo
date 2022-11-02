package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/inuoshios/keepinfo/internal/models"
)

func (u *postgresDBRepo) CreateSession(session *models.Session) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	var newId []byte

	query := `
	INSERT INTO contacts (user_id, email, refresh_token, user_agent, client_ip, is_blocked, expired_at, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`

	err := u.DB.QueryRowContext(ctx, query,
		session.UserID,
		session.Email,
		session.RefreshToken,
		session.UserAgent,
		session.ClientIP,
		session.IsBlocked,
		session.ExpiredAt,
		session.CreatedAt,
	).Scan(&newId)

	if err != nil {
		return "", fmt.Errorf("error inserting into table %w", err)
	}

	return string(newId), nil
}
