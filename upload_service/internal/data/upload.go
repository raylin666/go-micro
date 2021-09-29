package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"upload_service/internal/biz"
)

type UploadRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUploadRepo(data *Data, logger log.Logger) biz.UploadRepo {
	return &UploadRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UploadRepo) StreamUploadFile(ctx context.Context, g *biz.Upload) error {
	return nil
}

func (r *UploadRepo) UrlUploadFile(ctx context.Context, g *biz.Upload) error {
	return nil
}