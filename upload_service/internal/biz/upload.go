package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/raylin666/go-micro-protoc/upload/v1"
	uuid "github.com/raylin666/go-micro-protoc/uuid/v1"
	"time"
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
func (uc *UploadUsecase) StreamUploadFile(ctx context.Context, g *Upload) (string, error) {
	grpcClient, err := pool.GetUuidGRPCClientPool()
	if err != nil {
		return "", err
	}

	generateUuid, err := grpcClient.Generate(ctx, &uuid.GenerateRequest{})
	if err != nil {
		return "", err
	}

	storagePathFile := fmt.Sprintf(
		"micro/%d/%02d%02d/%s",
		time.Now().Year(),
		int(time.Now().Month()),
		time.Now().Day(),
		generateUuid.GetValue(),
		)
	if len(g.StreamUploadFile.GetMimeType()) > 0 {
		storagePathFile = fmt.Sprintf("%s.%s", storagePathFile, g.StreamUploadFile.GetMimeType())
	}

	put, err := qiniu.Get().FormUploaderPut([]byte(g.StreamUploadFile.GetStream()), storagePathFile)
	if err != nil {
		return "", err
	}

	fmt.Println(put)

	return "put", nil
}

