package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/raylin666/go-micro-protoc/upload/v1"
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

func (uc *UploadUsecase) StreamUploadFile(ctx context.Context, g *Upload) (string, error) {
	storagePathFile := "x/ccc.jpg"
	put, err := qiniu.Get().FormUploaderPut([]byte(g.StreamUploadFile.GetStream()), storagePathFile)
	if err != nil {
		return "", err
	}

	fmt.Println(put)

	return "put", nil
}

