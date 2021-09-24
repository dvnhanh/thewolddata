package config

import (
	"github.com/dvnhanh/thewolddata/internal/shared"
	"github.com/dvnhanh/thewolddata/pkg/database"
	"github.com/dvnhanh/thewolddata/pkg/http"
	"github.com/spf13/viper"
)

func NewViperConfig(cfgFolderPath, cfgFileName string) (Config, error) {
	p := &viperConfig{
		cfgFolderPath: cfgFolderPath,
		cfgFileName:   cfgFileName,
	}
	err := p.setup()
	return p, err
}

type viperConfig struct {
	cfgFolderPath string
	cfgFileName   string
}

func (p *viperConfig) setup() error {
	viper.SetConfigName(p.cfgFileName)
	viper.AddConfigPath(p.cfgFolderPath)
	viper.AutomaticEnv()

	// Read in config file
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.WatchConfig()
	return nil
}

func (p *viperConfig) GetDBConfig() *database.Config {
	return &database.Config{
		UserName:        viper.GetString(shared.CONFIG_KEY_MYSQL_USERNAME),
		Password:        viper.GetString(shared.CONFIG_KEY_MYSQL_PASSWORD),
		Address:         viper.GetString(shared.CONFIG_KEY_MYSQL_ADDRESS),
		Database:        viper.GetString(shared.CONFIG_KEY_MYSQL_DATABASE),
		NumberMaxConns:  viper.GetInt(shared.CONFIG_KEY_MYSQL_NUMBER_MAX_CONNS),
		NumberIdleConns: viper.GetInt(shared.CONFIG_KEY_MYSQL_NUMBER_IDLE_CONNS),
	}
}

func (p *viperConfig) GetHTTPConfig() *http.Config {
	return &http.Config{
		Address: viper.GetString(shared.CONFIG_KEY_HTTP_ADDRESS),
	}
}
