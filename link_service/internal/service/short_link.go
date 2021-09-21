package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"link_service/internal/biz"

	pb "github.com/raylin666/go-micro-protoc/link/v1"
)

type ShortLinkService struct {
	pb.UnimplementedShortLinkServer

	uc  *biz.ShortLinkUsecase
	log *log.Helper
}

func NewShortLinkService(uc *biz.ShortLinkUsecase, logger log.Logger) *ShortLinkService {
	return &ShortLinkService{
		uc: uc,
		log: log.NewHelper(logger),
	}
}

func (s *ShortLinkService) GenerateShortLink(ctx context.Context, req *pb.GenerateShortLinkRequest) (*pb.GenerateShortLinkReply, error) {
	url, err := s.uc.GenerateShortLink(ctx, &biz.ShortLink{
		GenerateShortLink: req,
	})

	return &pb.GenerateShortLinkReply{
		Url: url,
	}, err
}
