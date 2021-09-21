package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"uuid_service/internal/biz"
)

type uuidRepo struct {
	data *Data
	log  *log.Helper
}

// NewUuidRepo .
func NewUuidRepo(data *Data, logger log.Logger) biz.UuidRepo {
	return &uuidRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *uuidRepo) GenerateUuid(ctx context.Context, g *biz.Uuid) error {
	return nil
}