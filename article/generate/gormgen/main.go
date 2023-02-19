package main

import (
	"flag"
	kratos_config "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	utils_logger "github.com/raylin666/go-utils/logger"
	"github.com/raylin666/go-utils/server/system"
	"mt/config"
	genDb "mt/generate/gormgen/db"
	"mt/internal/app"
	"mt/pkg/db"
	"mt/pkg/logger"
	"mt/pkg/repositories"
)

var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../.env.yaml", "config path, eg: -conf .env.yaml")
}

func main() {
	c := kratos_config.New(
		kratos_config.WithSource(
			file.NewSource(flagconf),
		),
	)

	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc config.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// 初始化 Datetime
	app.Datetime = system.NewDatetime(
		system.WithLocation(bc.Datetime.Location),
		system.WithCSTLayout(bc.Datetime.CstLayout))

	log, err := logger.NewJSONLogger(
		utils_logger.WithField(utils_logger.AppKey, bc.App.Name),
		utils_logger.WithField(utils_logger.EnvironmentKey, bc.Environment),
		utils_logger.WithTimeLayout(app.Datetime.CSTLayout()))
	if err != nil {
		panic(err)
	}

	defaultDB, err := db.NewDb(repositories.DB_CONNECTION_DEFAULT_NAME, bc.Data.Database.Default, log)
	if err != nil {
		panic(err)
	}

	// 生成文件存放目录
	var outPath = "../../internal/repositories/dbrepo/query"

	// 执行生成默认数据库对应的模型文件
	genDb.NewGeneratorDefaultDb(defaultDB, outPath)
}
