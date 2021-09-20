package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ShortLink struct {

}

type ShortLinkRepo interface {
	GenerateShortLink(context.Context, *ShortLink) error
}

type ShortLinkUsecase struct {
	repo ShortLinkRepo
	log  *log.Helper
}

func NewShortLinkUsecase(repo ShortLinkRepo, logger log.Logger) *ShortLinkUsecase {
	return &ShortLinkUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ShortLinkUsecase) GenerateShortLink(ctx context.Context, g *ShortLink) error {
	return uc.repo.GenerateShortLink(ctx, g)
}