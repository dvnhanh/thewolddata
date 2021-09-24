package mysqlRepo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dvnhanh/thewolddata/internal/core/ports"
	"github.com/dvnhanh/thewolddata/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestMysqlMock(t *testing.T) {
	connector, mock, err := database.NewSqlmockConnector()
	assert.NoError(t, err)

	db := database.NewDatabase(connector)

	err = db.Connect(nil)
	assert.NoError(t, err)

	repo := NewMysqlRepo(db)

	t.Run("success_cases", func(t *testing.T) {
		testRegisterSuccessMock(t, mock, repo)
	})

	t.Run("failed_cases", func(t *testing.T) {
		testRegisterFailedMock(t, mock, repo)
	})

	err = db.Disconnect()
	assert.NoError(t, err)
}

func testRegisterSuccessMock(t *testing.T, mock sqlmock.Sqlmock, repo ports.ThewolddataMysqlRepoS) {

}

func testRegisterFailedMock(t *testing.T, mock sqlmock.Sqlmock, repo ports.ThewolddataMysqlRepoS) {

}
