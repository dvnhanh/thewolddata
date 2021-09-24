package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViperConfig(t *testing.T) {
	t.Run("error_cases", func(t *testing.T) {
		testViperErrorCase(t)
	})

	t.Run("success_cases", func(t *testing.T) {
		testViperSuccessCase(t)
	})
}

func testViperErrorCase(t *testing.T) {
	var err error
	_, err = NewViperConfig("../../../config", "config.file")
	assert.Error(t, err)
	_, err = NewViperConfig("config.folder", "config.local.test")
	assert.Error(t, err)
	_, err = NewViperConfig("config.folder", "config.file")
	assert.Error(t, err)
}

func testViperSuccessCase(t *testing.T) {
	cfg, err := NewViperConfig("../../../config", "config.test")
	assert.NoError(t, err)

	assert.Equal(t, cfg.GetDBConfig().Address, "127.0.0.1:3306")
	assert.Equal(t, cfg.GetDBConfig().UserName, "root")
	assert.Equal(t, cfg.GetDBConfig().Password, "root")
	assert.Equal(t, cfg.GetDBConfig().Database, "theworlddata")
	assert.Equal(t, cfg.GetDBConfig().NumberIdleConns, 2)
	assert.Equal(t, cfg.GetDBConfig().NumberMaxConns, 100)

	assert.Equal(t, cfg.GetHTTPConfig().Address, "8000")
}
