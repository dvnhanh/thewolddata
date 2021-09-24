package mysqlRepo

import (
	"context"

	"github.com/dvnhanh/thewolddata/internal/core/port"
	"github.com/dvnhanh/thewolddata/pkg/database"
)

func NewMysqlRepo(db database.Database) port.TheworlddataMysqlRepoS {
	return &mysqlRepo{
		db: db,
	}
}

type mysqlRepo struct {
	db database.Database
}

func (repo *mysqlRepo) Register(email, password string) error {
	return repo.db.Transaction(context.Background(), func(ctx context.Context, conn database.Connection) error {
		_, err := conn.ExecContext(
			ctx,
			"CALL `theworlddata`.resigter_account(?, ?)",
			email,
			password,
		)
		return err
	})
}
