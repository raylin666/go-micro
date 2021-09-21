package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/raylin666/go-micro-protoc/uuid/v1"
)

type Uuid struct {
	GenerateUuid *pb.GenerateUuidRequest
}

type UuidRepo interface {
	GenerateUuid(context.Context, *Uuid) (string, error)
}

type UuidUsecase struct {
	repo UuidRepo
	log  *log.Helper
}

func NewUuidUsecase(repo UuidRepo, logger log.Logger) *UuidUsecase {
	return &UuidUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UuidUsecase) GenerateUuid(ctx context.Context, g *Uuid) (string, error) {
	return uc.repo.GenerateUuid(ctx, g)
}
