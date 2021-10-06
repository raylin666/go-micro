package service

import (
	repositorie_casbin "auth_service/repositorie/casbin"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/raylin666/go-micro-protoc/errors"
	pb "github.com/raylin666/go-micro-protoc/services/auth/v1"
	"google.golang.org/protobuf/types/known/emptypb"

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

	return &pb.GetRolesForUserReply{Roles: roles}, nil
}

// AddRoleForUser 用户添加角色
func (s *AuthService) AddRoleForUser(ctx context.Context, req *pb.AddRoleForUserRequest) (*pb.AddRoleForUserReply, error) {
	is_ok, err := s.uc.AddRoleForUser(ctx, &biz.Auth{
		AddRoleForUser: req,
	})

	if err != nil {
		return nil, err
	}

	return &pb.AddRoleForUserReply{IsOk: is_ok}, nil
}

// GetUsersForRole 获取具有角色的用户
func (s *AuthService) GetUsersForRole(ctx context.Context, req *pb.GetUsersForRoleRequest) (*pb.GetUsersForRoleReply, error) {
	users, err := s.uc.GetUsersForRole(ctx, &biz.Auth{
		GetUsersForRole: req,
	})

	if err != nil {
		return nil, err
	}

	return &pb.GetUsersForRoleReply{Users: users}, nil
}

// HasRoleForUser 确定用户是否具有角色
func (s *AuthService) HasRoleForUser(ctx context.Context, req *pb.HasRoleForUserRequest) (*pb.HasRoleForUserReply, error) {
	has, err := s.uc.HasRoleForUser(ctx, &biz.Auth{
		HasRoleForUser: req,
	})

	if err != nil {
		return nil, err
	}

	return &pb.HasRoleForUserReply{Has: has}, nil
}

// DeleteRoleForUser 删除用户的角色
func (s *AuthService) DeleteRoleForUser(ctx context.Context, req *pb.DeleteRoleForUserRequest) (*emptypb.Empty, error) {
	is_ok, err := s.uc.DeleteRoleForUser(ctx, &biz.Auth{
		DeleteRoleForUser: req,
	})

	if err != nil {
		return nil, err
	}

	if !is_ok {
		return nil, errors.ErrorResourceDeleteError("resource delete error")
	}

	return nil, nil
}

// DeleteRolesForUser  删除用户的所有角色
func (s *AuthService) DeleteRolesForUser(ctx context.Context, req *pb.DeleteRolesForUserRequest) (*emptypb.Empty, error) {
	is_ok, err := s.uc.DeleteRolesForUser(ctx, &biz.Auth{
		DeleteRolesForUser: req,
	})

	if err != nil {
		return nil, err
	}

	if !is_ok {
		return nil, errors.ErrorResourceDeleteError("resource delete error")
	}

	return nil, nil
}
