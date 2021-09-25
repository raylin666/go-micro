package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb_link "github.com/raylin666/go-micro-protoc/link/v1"
	uuid "github.com/raylin666/go-micro-protoc/uuid/v1"
	"link_service/internal/conf"
	"link_service/repositorie/pool"
	"link_service/util/binary"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type ShortLink struct {
	GenerateShortLink *pb_link.GenerateShortUrlRequest
	ShortUrlToLongUrl *pb_link.TransformLongUrlRequest

	Ident   string
	LongUrl string
}

type ShortLinkRepo interface {
	GenerateShortUrl(context.Context, *ShortLink) error
	TransformLongUrl(context.Context, *ShortLink) (string, error)
}

type ShortLinkUsecase struct {
	repo ShortLinkRepo
	log  *log.Helper
}

func NewShortLinkUsecase(repo ShortLinkRepo, logger log.Logger) *ShortLinkUsecase {
	return &ShortLinkUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *ShortLinkUsecase) GenerateShortUrl(ctx context.Context, g *ShortLink) (string, error) {
	grpcClient, err := pool.GetUuidGRPCClientPool()
	if err != nil {
		return "", err
	}

	generateUuid, err := grpcClient.Generate(ctx, &uuid.GenerateRequest{Type: "time_rand"})
	if err != nil {
		return "", err
	}

	binaryTransform := binary.NewBinaryTransform()
	rand.Seed(time.Now().UnixNano())
	transformInt, _ := strconv.Atoi(fmt.Sprintf("%s%d", generateUuid.GetValue(), rand.Intn(999999 - 100000) + 100000))
	g.LongUrl = g.GenerateShortLink.GetUrl()
	g.Ident = binaryTransform.DecToB64(transformInt)
	url := conf.GetStore().GetApp().GetLinkDomain() + g.Ident

	err = uc.repo.GenerateShortUrl(ctx, g)
	if err != nil {
		return "", err
	}

	return url, err
}

func (uc *ShortLinkUsecase) TransformLongUrl(ctx context.Context, g *ShortLink) (string, error) {
	g.Ident = strings.Replace(g.ShortUrlToLongUrl.GetUrl(), conf.GetStore().GetApp().GetLinkDomain(), "", 1)
	url, err := uc.repo.TransformLongUrl(ctx, g)
	if err != nil {
		return "", err
	}

	return url, nil
}