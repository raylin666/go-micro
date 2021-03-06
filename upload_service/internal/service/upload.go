package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/raylin666/go-micro-protoc/services/upload/v1"
	"upload_service/internal/biz"
)

// UploadService is a greeter service.
type UploadService struct {
	pb.UnimplementedUploadServer

	uc  *biz.UploadUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewUploadService(uc *biz.UploadUsecase, logger log.Logger) *UploadService {
	return &UploadService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UploadService) StreamUploadFile(ctx context.Context, req *pb.StreamUploadFileRequest) (*pb.StreamUploadFileReply, error) {
	ret, err := s.uc.StreamUploadFile(ctx, &biz.Upload{
		StreamUploadFile: req,
	})

	if err != nil {
		return nil, err
	}

	return &pb.StreamUploadFileReply{
		Hash: ret.Hash,
		Key: ret.Key,
		Fsize: ret.Fsize,
		Url: ret.Url,
		Name: ret.Name,
		Bucket: ret.Bucket,
		Ext: ret.Ext,
		MimeType: ret.MimeType,
		Uuid: ret.Uuid,
	}, nil
}

func (s *UploadService) UrlUploadFile(ctx context.Context, req *pb.UrlUploadFileRequest) (*pb.UrlUploadFileReply, error) {
	ret, err := s.uc.UrlUploadFile(ctx, &biz.Upload{
		UrlUploadFile: req,
	})

	if err != nil {
		return nil, err
	}

	return &pb.UrlUploadFileReply{
		Hash: ret.Hash,
		Key: ret.Key,
		Fsize: ret.Fsize,
		Url: ret.Url,
		MimeType: ret.MimeType,
	}, nil
}

