package biz

import (
	repositorie_casbin "auth_service/internal/repositorie/casbin"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/raylin666/go-micro-protoc/services/auth/v1"
)

type Auth struct {
	GetRolesForUser *pb.GetRolesForUserRequest
}

type AuthRepo interface {
	GetRolesForUser(context.Context, *Auth) error
}

type AuthUsecase struct {
	repo AuthRepo
	log  *log.Helper
	casbin *casbin.Enforcer
}

func NewAuthUsecase(repo AuthRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: log.NewHelper(logger), casbin: repositorie_casbin.GetEnforcer()}
}

func (uc *AuthUsecase) GetRolesForUser(ctx context.Context, g *Auth) ([]string, error) {
	var (
		roles []string
		err error
	)

	roles, err = uc.casbin.GetRolesForUser(g.GetRolesForUser.GetUser())
	if err != nil {
		return roles, err
	}

	_ = uc.repo.GetRolesForUser(ctx, g)

	return roles, err
}

