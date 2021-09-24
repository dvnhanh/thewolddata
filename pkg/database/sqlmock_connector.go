package database

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewSqlmockConnector() (Connector, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	return &sqlmockConnector{
		db:    db,
		ready: false,
	}, mock, nil
}

type sqlmockConnector struct {
	db    *sql.DB
	ready bool
}

func (p *sqlmockConnector) Open(cfg *Config) error {
	p.ready = true
	return nil
}

func (p *sqlmockConnector) Close() error {
	if !p.ready {
		return nil
	}
	return p.db.Close()
}

func (p *sqlmockConnector) GetDB() *sql.DB {
	if !p.ready {
		return nil
	}
	return p.db
}
