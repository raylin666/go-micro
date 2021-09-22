package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"link_service/internal/biz"
	"link_service/internal/data/model"
)

type shortLinkRepo struct {
	data *Data
	log  *log.Helper
}

// NewShortLinkRepo .
func NewShortLinkRepo(data *Data, logger log.Logger) biz.ShortLinkRepo {
	return &shortLinkRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *shortLinkRepo) GenerateShortLink(ctx context.Context, g *biz.ShortLink) error {
	return model.LinkRelation{}.Create(r.data.db, g.Ident, g.LongUrl)
}
