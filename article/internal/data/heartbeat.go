package data

import (
	"context"
	"mt/pkg/logger"

	"mt/internal/biz"
)

type heartbeatRepo struct {
	data *Data
	log  *logger.Logger
}

func NewHeartbeatRepo(data *Data, logger *logger.Logger) biz.HeartbeatRepo {
	return &heartbeatRepo{
		data: data,
		log:  logger,
	}
}

func (r *heartbeatRepo) PONE(ctx context.Context) error {
	return nil
}
