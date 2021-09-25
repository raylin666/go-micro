package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/raylin666/go-micro-protoc/link/v1"
	"link_service/internal/biz"
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

// GenerateShortUrl 生成短链接 /**
func (s *ShortLinkService) GenerateShortUrl(ctx context.Context, req *pb.GenerateShortUrlRequest) (*pb.GenerateShortUrlReply, error) {
	url, err := s.uc.GenerateShortUrl(ctx, &biz.ShortLink{
		GenerateShortLink: req,
	})

	return &pb.GenerateShortUrlReply{
		Url: url,
	}, err
}

// TransformLongUrl 通过短链接拉取长链接 /**
func (s *ShortLinkService) TransformLongUrl(ctx context.Context, req * pb.TransformLongUrlRequest) (*pb.TransformLongUrlReply, error) {
	url, err := s.uc.TransformLongUrl(ctx, &biz.ShortLink{
		ShortUrlToLongUrl: req,
	})

	return &pb.TransformLongUrlReply{
		Url: url,
	}, err
}
