package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"mt/api/v1"

	"mt/internal/biz"
)

type HeartbeatService struct {
	v1.UnimplementedHeartbeatServer

	uc *biz.HeartbeatUsecase
}

func NewHeartbeatService(uc *biz.HeartbeatUsecase) *HeartbeatService {
	return &HeartbeatService{uc: uc}
}

func (s *HeartbeatService) PONE(ctx context.Context, req *emptypb.Empty) (*v1.PONEResponse, error) {
	err := s.uc.PONE(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.PONEResponse{Message: "PONE"}, nil
}
