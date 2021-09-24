package config

import (
	"github.com/dvnhanh/thewolddata/pkg/database"
	"github.com/dvnhanh/thewolddata/pkg/http"
)

type Config interface {
	GetDBConfig() *database.Config
	GetHTTPConfig() *http.Config
}
