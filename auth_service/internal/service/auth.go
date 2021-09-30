package service

import (
	repositorie_casbin "auth_service/internal/repositorie/casbin"
	"context"
	"github.com/casbin/casbin/v2"
	pb "github.com/raylin666/go-micro-protoc/services/auth/v1"

	"auth_service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// AuthService is a greeter service.
type AuthService struct {
	pb.UnimplementedAuthServer

	uc  *biz.AuthUsecase
	log *log.Helper
	casbin *casbin.Enforcer
}

// NewAuthService new a greeter service.
func NewAuthService(uc *biz.AuthUsecase, logger log.Logger) *AuthService {
	return &AuthService{uc: uc, log: log.NewHelper(logger), casbin: repositorie_casbin.GetEnforcer()}
}

// GetRolesForUser 获取用户具有的角色
func (s *AuthService) GetRolesForUser(ctx context.Context, req *pb.GetRolesForUserRequest) (*pb.GetRolesForUserReply, error) {
	roles, err := s.uc.GetRolesForUser(ctx, &biz.Auth{
		GetRolesForUser: req,
	})

	if err != nil {
		return nil, err
	}

	return &pb.GetRolesForUserReply{
		Roles: roles,
	}, nil
}
