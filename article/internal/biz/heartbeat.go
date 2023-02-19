package biz

import (
	"context"
	"mt/pkg/logger"
)

type Heartbeat struct {
}

type HeartbeatRepo interface {
	PONE(context.Context) error
}

type HeartbeatUsecase struct {
	repo HeartbeatRepo
	log  *logger.Logger
}

func NewHeartbeatUsecase(repo HeartbeatRepo, logger *logger.Logger) *HeartbeatUsecase {
	return &HeartbeatUsecase{repo: repo, log: logger}
}

func (uc *HeartbeatUsecase) PONE(ctx context.Context) error {
	return uc.repo.PONE(ctx)
}
