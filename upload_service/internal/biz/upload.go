package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mitchellh/mapstructure"
	pb "github.com/raylin666/go-micro-protoc/services/upload/v1"
	uuid "github.com/raylin666/go-micro-protoc/services/uuid/v1"
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

