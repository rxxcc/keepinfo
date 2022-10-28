package dbrepo

import (
	"database/sql"

	"github.com/inuoshios/keepinfo/internal/config"
	"github.com/inuoshios/keepinfo/internal/repository"
)

type postgresDBRepo struct {
	App *config.Config
	DB  *sql.DB
}

func NewPostgresRepo(a *config.Config, conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
