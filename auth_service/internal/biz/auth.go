package biz

import (
	repositorie_casbin "auth_service/repositorie/casbin"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/raylin666/go-micro-protoc/services/auth/v1"
)

type Auth struct {
	GetRolesForUser    *pb.GetRolesForUserRequest
	AddRoleForUser     *pb.AddRoleForUserRequest
	GetUsersForRole    *pb.GetUsersForRoleRequest
	HasRoleForUser     *pb.HasRoleForUserRequest
	DeleteRoleForUser  *pb.DeleteRoleForUserRequest
	DeleteRolesForUser *pb.DeleteRolesForUserRequest
	DeleteUser         *pb.DeleteUserRequest
	DeleteRole         *pb.DeleteRoleRequest
}

type AuthRepo interface {
	GetRolesForUser(context.Context, *Auth) error
	AddRoleForUser(context.Context, *Auth) error
	GetUsersForRole(context.Context, *Auth) error
	HasRoleForUser(context.Context, *Auth) error
	DeleteRoleForUser(context.Context, *Auth) error
	DeleteRolesForUser(context.Context, *Auth) error
	DeleteUser(context.Context, *Auth) error
	DeleteRole(context.Context, *Auth) error
}

type AuthUsecase struct {
	repo   AuthRepo
	log    *log.Helper
	casbin *casbin.Enforcer
}

func NewAuthUsecase(repo AuthRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: log.NewHelper(logger), casbin: repositorie_casbin.GetEnforcer()}
}

func (uc *AuthUsecase) GetRolesForUser(ctx context.Context, g *Auth) ([]string, error) {
	var (
		roles []string
		err   error
	)

	roles, err = uc.casbin.GetRolesForUser(g.GetRolesForUser.GetUser())
	if err != nil {
		return roles, err
	}

	_ = uc.repo.GetRolesForUser(ctx, g)

	return roles, err
}

func (uc *AuthUsecase) AddRoleForUser(ctx context.Context, g *Auth) (bool, error) {
	var (
		is_ok bool
		err   error
	)

	is_ok, err = uc.casbin.AddRoleForUser(g.AddRoleForUser.GetUser(), g.AddRoleForUser.GetRole())
	if err != nil {
		return is_ok, err
	}

	_ = uc.repo.AddRoleForUser(ctx, g)

	return is_ok, err
}

func (uc *AuthUsecase) GetUsersForRole(ctx context.Context, g *Auth) ([]string, error) {
	var (
		users []string
		err   error
	)

	users, err = uc.casbin.GetUsersForRole(g.GetUsersForRole.GetRole())
	if err != nil {
		return users, err
	}

	_ = uc.repo.GetUsersForRole(ctx, g)

	return users, err
}

func (uc *AuthUsecase) HasRoleForUser(ctx context.Context, g *Auth) (bool, error) {
	var (
		has bool
		err error
	)

	has, err = uc.casbin.HasRoleForUser(g.HasRoleForUser.GetUser(), g.HasRoleForUser.GetRole())
	if err != nil {
		return has, err
	}

	_ = uc.repo.GetUsersForRole(ctx, g)

	return has, err
}

func (uc *AuthUsecase) DeleteRoleForUser(ctx context.Context, g *Auth) (bool, error) {
	var (
		is_ok bool
		err error
	)

	is_ok, err = uc.casbin.DeleteRoleForUser(g.DeleteRoleForUser.GetUser(), g.DeleteRoleForUser.GetRole())
	if err != nil {
		return is_ok, err
	}

	_ = uc.repo.DeleteRoleForUser(ctx, g)

	return is_ok, err
}

func (uc *AuthUsecase) DeleteRolesForUser(ctx context.Context, g *Auth) (bool, error) {
	var (
		is_ok bool
		err error
	)

	is_ok, err = uc.casbin.DeleteRolesForUser(g.DeleteRolesForUser.GetUser())
	if err != nil {
		return is_ok, err
	}

	_ = uc.repo.DeleteRolesForUser(ctx, g)

	return is_ok, err
}
