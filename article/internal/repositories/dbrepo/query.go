package dbrepo

import (
	"gorm.io/gorm"
	"mt/internal/repositories/dbrepo/query"
	"mt/pkg/repositories"
)

// NewDefaultDb 创建默认数据库
func NewDefaultDb(dbRepo repositories.DbRepo) *gorm.DB {
	return dbRepo.DB(repositories.DB_CONNECTION_DEFAULT_NAME).Get().DB()
}

// NewDefaultDbQuery 创建默认数据库查询
func NewDefaultDbQuery(dbRepo repositories.DbRepo) *query.Query {
	return query.Use(NewDefaultDb(dbRepo))
}


