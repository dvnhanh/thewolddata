package database

import (
	"context"
	"database/sql"
	"errors"
)

// Container methob have purpose convertation with databse

type Connection interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
}

type CommandFunc func(context.Context, Connection) error

type Database interface {
	Connect(cfg *Config) error
	Disconnect() error
	Transaction(ctx context.Context, conn ...CommandFunc) error
	WithoutTransaction(ctx context.Context, conn ...CommandFunc) error
}

func NewDatabase(connector Connector) Database {
	return &database{
		connector: connector,
	}
}

type database struct {
	connector Connector
}

func (p *database) Connect(cfg *Config) error {
	return p.connector.Open(cfg)
}

func (p *database) Disconnect() error {
	return p.connector.Close()
}

func (p *database) Transaction(ctx context.Context, cmdFuncs ...CommandFunc) error {
	db := p.connector.GetDB()
	if db == nil {
		return errors.New("forgot connect databse")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, cmdFunc := range cmdFuncs {
		err = cmdFunc(ctx, tx)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}
	}

	return tx.Commit()
}

func (p *database) WithoutTransaction(ctx context.Context, cmdFuncs ...CommandFunc) error {
	db := p.connector.GetDB()
	if db == nil {
		return errors.New("forgot connect database")
	}

	for _, cmdFunc := range cmdFuncs {
		err := cmdFunc(ctx, db)
		if err != nil {
			return err
		}
	}

	return nil
}
