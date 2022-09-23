package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// DB is the Database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const (
	maxOpenedConn = 10
	maxIdleConn   = 5
	maxIdleTime   = 5 * time.Minute
	maxLifetime   = 1 * time.Hour
)

// ConnectSQL created Database pool for PG
func ConnectSQL(dsn string) (*DB, error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(maxOpenedConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxIdleTime(maxIdleTime)
	db.SetConnMaxLifetime(maxLifetime)

	dbConn.SQL = db

	if err = testDB(db); err != nil {
		return nil, err
	}

	return dbConn, nil
}

// testDB helps pings the Database
func testDB(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}

// NewDatabase create a new database for the project
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
