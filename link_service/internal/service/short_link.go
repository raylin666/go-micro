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

/**
	生成短链接
 */
func (s *ShortLinkService) GenerateShortLink(ctx context.Context, req *pb.GenerateShortLinkRequest) (*pb.GenerateShortLinkReply, error) {
	url, err := s.uc.GenerateShortLink(ctx, &biz.ShortLink{
		GenerateShortLink: req,
	})

	return &pb.GenerateShortLinkReply{
		Url: url,
	}, err
}

/**
	通过短链接拉取长链接
 */
func (s *ShortLinkService) ShortUrlToLongUrl(ctx context.Context, req * pb.ShortUrlToLongUrlRequest) (*pb.ShortUrlToLongUrlReply, error) {
	url, err := s.uc.ShortUrlToLongUrl(ctx, &biz.ShortLink{
		ShortUrlToLongUrl: req,
	})

	return &pb.ShortUrlToLongUrlReply{
		Url: url,
	}, err
}
