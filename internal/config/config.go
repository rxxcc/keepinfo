package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/inuoshios/keepinfo/internal/database"
)

// Config consists of all the application config
type Config struct {
	SQL     *database.DB
	Session *scs.SessionManager
}
