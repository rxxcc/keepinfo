package postgres

import (
	"database/sql"

	"github.com/ixxiv/keepinfo/internal/config"
	"github.com/ixxiv/keepinfo/internal/repository"
)

type postgres struct {
	App *config.Config
	DB  *sql.DB
}

// NewPostgresRepo initializes the DatabaseRepo
func NewPostgresRepo(a *config.Config, conn *sql.DB) repository.DatabaseRepo {
	return &postgres{
		App: a,
		DB:  conn,
	}
}
