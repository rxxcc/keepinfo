package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/inuoshios/keepinfo/internal/models"
)

func (u *postgres) CreateSession(session *models.Session) (models.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	var newSession models.Session

	query := `
	INSERT INTO sessions (id, user_id, refresh_token, user_agent, client_ip, is_blocked, expired_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, user_id, refresh_token, user_agent, client_ip, is_blocked, expired_at`

	err := u.DB.QueryRowContext(ctx, query,
		session.ID,
		session.UserID,
		session.RefreshToken,
		session.UserAgent,
		session.ClientIP,
		session.IsBlocked,
		session.ExpiredAt,
	).Scan(
		&newSession.ID,
		&newSession.UserID,
		&newSession.RefreshToken,
		&newSession.UserAgent,
		&newSession.ClientIP,
		&newSession.IsBlocked,
		&newSession.ExpiredAt,
	)

	if err != nil {
		return newSession, fmt.Errorf("error inserting into sessions table %w", err)
	}

	return newSession, nil
}

func (u *postgres) GetSession(id uuid.UUID) (models.Session, error) {
	var session models.Session
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	query := `
	SELECT id, user_id, refresh_token, user_agent, client_ip, is_blocked, expired_at, created_at
	FROM sessions WHERE id = $1 LIMIT 1`

	rows := u.DB.QueryRowContext(ctx, query, id)
	err := rows.Scan(
		&session.ID,
		&session.UserID,
		&session.RefreshToken,
		&session.UserAgent,
		&session.ClientIP,
		&session.IsBlocked,
		&session.ExpiredAt,
		&session.CreatedAt,
	)

	if err != nil {
		return session, fmt.Errorf("error getting sessions %w", err)
	}

	return session, nil
}
