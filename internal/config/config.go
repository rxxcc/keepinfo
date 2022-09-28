package config

import (
	"log"
)

// Config consists of all the application config
type Config struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
