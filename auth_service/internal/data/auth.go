package data

import (
	"auth_service/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &AuthRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *AuthRepo) GetRolesForUser(ctx context.Context, g *biz.Auth) error {
	return nil
}
