package config

import (
	"log"

	"github.com/inuoshios/keepinfo/internal/database"
)

// Config consists of all the application config
type Config struct {
	SQL      *database.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
