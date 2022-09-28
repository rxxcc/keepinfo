package config

import (
	"database/sql"
	"log"
)

// Config consists of all the application config
type Config struct {
	SQL      *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
