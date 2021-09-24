package database

import (
	"database/sql"
	"fmt"
)

func NewMysqlConnector() Connector {
	return new(mysqlConnector)
}

type mysqlConnector struct {
	db *sql.DB
}

func (p *mysqlConnector) Open(cfg *Config) error {
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.UserName,
		cfg.Password,
		cfg.Address,
		cfg.Database,
	)
	db, err := sql.Open("mysql", str)
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(cfg.NumberIdleConns)
	db.SetMaxOpenConns(cfg.NumberMaxConns)
	db.SetConnMaxLifetime(0)
	p.db = db

	return nil
}

func (p *mysqlConnector) Close() error {
	return p.db.Close()
}

func (p *mysqlConnector) GetDB() *sql.DB {
	return p.db
}
