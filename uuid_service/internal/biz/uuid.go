package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/raylin666/go-micro-protoc/uuid/v1"
	uuid "github.com/satori/go.uuid"
	"time"
	"uuid_service/internal/constant"
)

type Uuid struct {
	GenerateUuid *pb.GenerateUuidRequest
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

func (uc *UuidUsecase) GenerateUuid(ctx context.Context, g *Uuid) (string, error) {
	var (
		value string
		err error
	)

	switch g.GenerateUuid.GetType() {
	case constant.UUID_TYPE_TIME_RAND:
		value = fmt.Sprintf("%v", time.Now().UnixNano() / 1e6)
		break
	default:
		value = uuid.NewV4().String()
	}

	err = uc.repo.GenerateUuid(ctx, g)

	return value, err
}
