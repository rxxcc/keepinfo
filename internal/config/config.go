package config

import "log"

type Config struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
