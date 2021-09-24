package mysqlRepo

import (
	"testing"

	"github.com/dvnhanh/thewolddata/internal/shared/config"
	"github.com/dvnhanh/thewolddata/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestMysqlIntegration(t *testing.T) {
	cfg, err := config.NewViperConfig("./../../../config", "config.test")
	assert.NoError(t, err)

	db := database.NewDatabase(
		database.NewDockerConnector([]string{
			"D:/pro-wct/my-app/theworlddata/api/sql/tables.sql",
			"D:/pro-wct/my-app/theworlddata/api/sql/procedures.sql",
		}))
	err = db.Connect(cfg.GetDBConfig())
	assert.NoError(t, err)

	t.Run("success_cases", func(t *testing.T) {
		testRegisterSuccessIntegration(t)
	})

	t.Run("failed_cases", func(t *testing.T) {
		testRegisterFailedIntegration(t)
	})

	err = db.Disconnect()
	assert.NoError(t, err)
}

func testRegisterSuccessIntegration(t *testing.T) {
	a := 2
	assert.Equal(t, a, 2)
}

func testRegisterFailedIntegration(t *testing.T) {
	a := 2
	assert.NotEqual(t, a, 5)
}
