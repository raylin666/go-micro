package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mitchellh/mapstructure"
	pb "github.com/raylin666/go-micro-protoc/upload/v1"
	uuid "github.com/raylin666/go-micro-protoc/uuid/v1"
	"mime"
	"time"
	"upload_service/internal/conf"
	"upload_service/repositorie/pool"
	"upload_service/repositorie/upload/qiniu"
)

type Upload struct {
	StreamUploadFile *pb.StreamUploadFileRequest
}

type UploadRepo interface {
	StreamUploadFile(context.Context, *Upload) error
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
	var (
		ret = new(qiniu.UploadPutRet)
	)

	grpcClient, err := pool.GetUuidGRPCClientPool()
	if err != nil {
		return nil, err
	}

	generateUuid, err := grpcClient.Generate(ctx, &uuid.GenerateRequest{})
	if err != nil {
		return nil, err
	}

	if len(g.StreamUploadFile.GetStream()) <= 0 {
		err = pb.ErrorStreamEmpty("stream is empty")
		return nil, err
	}

	storagePathFile := fmt.Sprintf(
		"micro/%d/%02d%02d/%s",
		time.Now().Year(),
		int(time.Now().Month()),
		time.Now().Day(),
		generateUuid.GetValue(),
		)
	if len(g.StreamUploadFile.GetMimeType()) > 0 {
		if ext, err := mime.ExtensionsByType(g.StreamUploadFile.GetMimeType()); err == nil {
			storagePathFile = fmt.Sprintf("%s%s", storagePathFile, ext[0])
		}
	}

	put, err := qiniu.Get().FormUploaderPut(g.StreamUploadFile.GetStream(), storagePathFile)
	if err != nil {
		return nil, err
	}

	// interface/map 转换 struct
	if err = mapstructure.Decode(put, &ret); err == nil {
		ret.Url = conf.GetStore().GetUpload().GetCdn() + ret.Key
	}

	return ret, nil
}

