package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Uuid struct {

}

type UuidRepo interface {
	GenerateUuid(context.Context, *Uuid) error
}

type UuidUsecase struct {
	repo UuidRepo
	log  *log.Helper
}

func NewUuidUsecase(repo UuidRepo, logger log.Logger) *UuidUsecase {
	return &UuidUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UuidUsecase) GenerateUuid(ctx context.Context, g *Uuid) error {
	return uc.repo.GenerateUuid(ctx, g)
}
