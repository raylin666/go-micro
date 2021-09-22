package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb_link "github.com/raylin666/go-micro-protoc/link/v1"
	uuid "github.com/raylin666/go-micro-protoc/uuid/v1"
	"link_service/internal/constant"
	"link_service/internal/util/binary"
	"link_service/internal/util/grpc"
	"math/rand"
	"strconv"
	"time"
)

type ShortLink struct {
	GenerateShortLink *pb_link.GenerateShortLinkRequest

	Ident   string
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
	rand.Seed(time.Now().UnixNano())
	transformInt, _ := strconv.Atoi(fmt.Sprintf("%s%d", generateUuid.GetValue(), rand.Intn(999999 - 100000) + 100000))
	g.LongUrl = g.GenerateShortLink.GetUrl()
	g.Ident = binaryTransform.DecToB64(transformInt)
	url := constant.LINK_DOMAIN + g.Ident

	err = uc.repo.GenerateShortLink(ctx, g)
	if err != nil {
		return "", err
	}

	return url, err
}
