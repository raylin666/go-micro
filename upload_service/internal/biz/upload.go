package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mitchellh/mapstructure"
	pb "github.com/raylin666/go-micro-protoc/services/upload/v1"
	uuid "github.com/raylin666/go-micro-protoc/services/uuid/v1"
	"mime"
	"net/url"
	"strings"
	"time"
	"upload_service/internal/conf"
	"upload_service/repositorie/pool"
	"upload_service/repositorie/upload/qiniu"
)

type Upload struct {
	StreamUploadFile *pb.StreamUploadFileRequest
	UrlUploadFile *pb.UrlUploadFileRequest
}

type UploadRepo interface {
	StreamUploadFile(context.Context, *Upload) error
	UrlUploadFile(context.Context, *Upload) error
}

type UploadUsecase struct {
	repo UploadRepo
	log  *log.Helper
}

func NewUploadUsecase(repo UploadRepo, logger log.Logger) *UploadUsecase {
	return &UploadUsecase{repo: repo, log: log.NewHelper(logger)}
}

// StreamUploadFile 数据流方式上传文件
func (uc *UploadUsecase) StreamUploadFile(ctx context.Context, g *Upload) (*qiniu.UploadPutRet, error) {
	grpcClient, err := pool.GetUuidGRPCClientPool()
	if err != nil {
		return nil, err
	}

	generateUuid, err := grpcClient.Generate(ctx, &uuid.GenerateRequest{})
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf(
		"micro/%d/%02d%02d/%s",
		time.Now().Year(),
		int(time.Now().Month()),
		time.Now().Day(),
		generateUuid.GetValue(),
		)
	if len(g.StreamUploadFile.GetMimeType()) > 0 {
		if ext, err := mime.ExtensionsByType(g.StreamUploadFile.GetMimeType()); err == nil {
			key = fmt.Sprintf("%s%s", key, ext[0])
		}
	}

	put, err := qiniu.Get().FormUploaderPut(g.StreamUploadFile.GetStream(), key)
	if err != nil {
		return nil, err
	}

	var (
		ret = new(qiniu.UploadPutRet)
	)

	// interface/map 转换 struct
	if err = mapstructure.Decode(put, &ret); err == nil {
		ret.Url = qiniu.Get().MakePublicURL(conf.GetStore().GetUpload().GetCdn(), ret.Key)
	}

	return ret, nil
}

// UrlUploadFile 通过URL资源地址方式上传文件
func (uc *UploadUsecase) UrlUploadFile(ctx context.Context, g *Upload) (*qiniu.UploadPutRet, error) {
	grpcClient, err := pool.GetUuidGRPCClientPool()
	if err != nil {
		return nil, err
	}

	generateUuid, err := grpcClient.Generate(ctx, &uuid.GenerateRequest{})
	if err != nil {
		return nil, err
	}

	ext := ""
	parse, _ := url.Parse(g.UrlUploadFile.GetUrl())
	path := strings.Split(parse.Path, "/")
	if len(path) > 1 {
		gpath := strings.Split(path[len(path) - 1], ".")
		if len(gpath) > 1 {
			ext = fmt.Sprintf(".%s", gpath[len(gpath) - 1])
		}
	}

	key := fmt.Sprintf(
		"micro/%d/%02d%02d/%s%s",
		time.Now().Year(),
		int(time.Now().Month()),
		time.Now().Day(),
		generateUuid.GetValue(),
		ext,
	)

	put, err := qiniu.Get().Fetch(g.UrlUploadFile.GetUrl(), key)
	if err != nil {
		return nil, err
	}

	var (
		ret = new(qiniu.UploadPutRet)
	)

	// interface/map 转换 struct
	if err = mapstructure.Decode(put, &ret); err == nil {
		ret.Url = qiniu.Get().MakePublicURL(conf.GetStore().GetUpload().GetCdn(), ret.Key)
	}

	return ret, nil
}
