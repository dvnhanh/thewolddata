package cmdTheworlddata

import (
	"github.com/dvnhanh/thewolddata/internal/core/ports"
	"github.com/dvnhanh/thewolddata/internal/core/services"
	"github.com/dvnhanh/thewolddata/internal/handler"
	"github.com/dvnhanh/thewolddata/internal/repos/mysqlRepo"
	"github.com/dvnhanh/thewolddata/internal/shared/config"
	"github.com/dvnhanh/thewolddata/pkg/database"
	"github.com/dvnhanh/thewolddata/pkg/log/logger"
	"github.com/dvnhanh/thewolddata/pkg/log/logger/entity"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

var (
	container     *dig.Container
	cfgFileName   string
	cfgFolderPath string
)

var RootCmd = &cobra.Command{
	Run: run,
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initialize)
	RootCmd.PersistentFlags().StringVar(&cfgFileName, "cfgFileName", "config.local", "config file name (defaul: config.local)")
	RootCmd.PersistentFlags().StringVar(&cfgFolderPath, "cfgFolderPath", "./config", "config path (default: ./config)")
}

func initialize() {
	container = buildContainer()
}

func run(cmd *cobra.Command, args []string) {
	container.Invoke(func(cfg config.Config, serverHandler handler.HTTPServer, logWriter logger.Logger) {
		// start server
		logWriter.Info(&entity.LogItem{
			Message: "start server the world data",
		})

		err := serverHandler.Begin(cfg.GetHTTPConfig().Address)
		if err != nil {
			logWriter.Error(&entity.LogItem{
				Message: "start server the world data FAILED",
			})
		}
	})
}

func buildContainer() *dig.Container {
	container := dig.New()

	//Setup log
	container.Provide(func() logger.Logger {
		return logger.NewConsoleLogger()
	})

	// Setup config
	container.Provide(func() config.Config {
		cfg, err := config.NewViperConfig(cfgFolderPath, cfgFileName)
		if err != nil {
			panic("Set up config failed!. Detail: " + err.Error())
		}
		return cfg
	})

	// Setup Repository
	container.Provide(func(cfg config.Config) ports.ThewolddataMysqlRepoS {
		db := database.NewDatabase(database.NewMysqlConnector())
		err := db.Connect(cfg.GetDBConfig())
		if err != nil {
			panic("Connect database failed!. Detail: " + err.Error())
		}

		return mysqlRepo.NewMysqlRepo(db)
	})

	// Setup Service
	container.Provide(func(repo ports.ThewolddataMysqlRepoS) ports.ThewolddataService {
		return services.NewTheWorldDataService(repo)
	})

	// Setup Handler
	container.Provide(func(svc ports.ThewolddataService) handler.HTTPServer {
		return handler.NewHTTPServer(svc)
	})

	return container
}
