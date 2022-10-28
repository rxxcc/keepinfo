package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sql.DB
}

const (
	maxOpenedConn = 10
	maxIdleConn   = 5
	maxIdleTime   = 5 * time.Minute
	maxLifetime   = 1 * time.Hour
)

// ConnectSQL created Database pool for PG
func ConnectSQL(dsn string) (*DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenedConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxIdleTime(maxIdleTime)
	db.SetConnMaxLifetime(maxLifetime)

	if err = testDB(db); err != nil {
		return nil, err
	}

	return &DB{
		db,
	}, nil
}

// testDB helps pings the Database
func testDB(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}
