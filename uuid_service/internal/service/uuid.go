package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"uuid_service/internal/biz"

	pb "github.com/raylin666/go-micro-protoc/uuid/v1"
)

type UuidService struct {
	pb.UnimplementedUuidServer

	uc  *biz.UuidUsecase
	log *log.Helper
}

func NewUuidService(uc *biz.UuidUsecase, logger log.Logger) *UuidService {
	return &UuidService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UuidService) GenerateUuid(ctx context.Context, req *pb.GenerateUuidRequest) (*pb.GenerateUuidReply, error) {
	value, err := s.uc.GenerateUuid(ctx, &biz.Uuid{
		GenerateUuid: req,
	})

	return &pb.GenerateUuidReply{
		Value: value,
	}, err
}
