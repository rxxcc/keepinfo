package config

import "github.com/alexedwards/scs/v2"

// Config consists of all the application config
type Config struct {
	Session *scs.SessionManager
}
