package db

import (
	"errors"
	"fmt"
	"github.com/raylin666/go-utils/db/gorm"
	gorm_db "gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
	"mt/config"
	"mt/pkg/logger"
	"time"
)

var _ Db = (*db)(nil)

type Predicate string

type Db interface {
	Get() gorm.Client
	Close() error
}

type db struct {
	client gorm.Client
	logger *logger.Logger
}

func NewDb(name string, config *config.DatabaseItem, logger *logger.Logger) (Db, error) {
	var rdb = new(db)
	rdb.logger = logger
	client, err := gorm.NewClient(
		gorm.WithDsn(config.Dsn),
		gorm.WithDriver(config.Driver),
		gorm.WithDbName(config.Dbname),
		gorm.WithHost(config.Host),
		gorm.WithUserName(config.Username),
		gorm.WithPassword(config.Password),
		gorm.WithCharset(config.Charset),
		gorm.WithPort(int(config.Port)),
		gorm.WithPrefix(config.Prefix),
		gorm.WithMaxIdleConn(int(config.MaxIdleConn)),
		gorm.WithMaxOpenConn(int(config.MaxOpenConn)),
		gorm.WithMaxLifeTime(time.Duration(config.MaxLifeTime)),
		gorm.WithParseTime(config.ParseTime),
		gorm.WithLoc(config.Loc))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("new db to %s db err", name))
	}

	rdb.client = client

	// 日志处理
	l := rdb.WithLogger()
	if l != nil {
		rdb.client.WithLogger(l)
	}
	// 插件处理
	_ = rdb.client.WithPluginBeforeHandler(rdb.BeforePluginHandler, rdb.AfterPluginHandler)

	return rdb, nil
}

func (db *db) Get() gorm.Client {
	return db.client
}

func (db *db) Close() error {
	return db.Get().SqlDB().Close()
}

// WithLogger DB 日志
func (db *db) WithLogger() gorm_logger.Interface {
	return NewLogger(
		db.logger,
		WithLoggerLevel(gorm_logger.Info),
		WithLoggerSlowThreshold(time.Second*1),
		WithLoggerIgnoreRecordNotFoundError(true))
}

// BeforePluginHandler DB 插件前置方法
func (db *db) BeforePluginHandler(rdb *gorm_db.DB) {}

// AfterPluginHandler DB 插件后置方法
func (db *db) AfterPluginHandler(rdb *gorm_db.DB, sql string, ts time.Time) {}
