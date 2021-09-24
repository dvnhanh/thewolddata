package database

import "database/sql"

// connection database

type Connector interface {
	Open(cfg *Config) error
	Close() error
	GetDB() *sql.DB
}
