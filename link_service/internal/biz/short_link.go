package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb_link "github.com/raylin666/go-micro-protoc/link/v1"
	uuid "github.com/raylin666/go-micro-protoc/uuid/v1"
	"link_service/internal/constant"
	"link_service/internal/util/binary"
	"link_service/internal/util/grpc"
	"strconv"
)

type ShortLink struct {
	GenerateShortLink *pb_link.GenerateShortLinkRequest

	Ident   int
	Value   string
	LongUrl string
}

type ShortLinkRepo interface {
	GenerateShortLink(context.Context, *ShortLink) error
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

func (uc *ShortLinkUsecase) GenerateShortLink(ctx context.Context, g *ShortLink) (string, error) {
	generateUuid, err := grpc.GRPCClientConn().GetUuidClient().GenerateUuid(ctx, &uuid.GenerateUuidRequest{
		Type: constant.UUID_TYPE_TIME_RAND,
	})

	if err != nil {
		return "", err
	}

	binaryTransform := binary.NewBinaryTransform()
	g.Ident, _ = strconv.Atoi(generateUuid.GetValue())
	g.LongUrl = g.GenerateShortLink.GetUrl()
	g.Value = binaryTransform.DecToB64(g.Ident)
	url := constant.LINK_DOMAIN + g.Value

	err = uc.repo.GenerateShortLink(ctx, g)
	if err != nil {
		return "", err
	}

	return url, err
}
