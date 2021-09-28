package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"uuid_service/internal/biz"

	pb "github.com/raylin666/go-micro-protoc/services/uuid/v1"
)

type UuidService struct {
	pb.UnimplementedUuidServer

	uc  *biz.UuidUsecase
	log *log.Helper
}

func NewUuidService(uc *biz.UuidUsecase, logger log.Logger) *UuidService {
	return &UuidService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UuidService) Generate(ctx context.Context, req *pb.GenerateRequest) (*pb.GenerateReply, error) {
	value, err := s.uc.Generate(ctx, &biz.Uuid{
		GenerateUuid: req,
	})

	return &pb.GenerateReply{
		Value: value,
	}, err
}
