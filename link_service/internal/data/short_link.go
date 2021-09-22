package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"link_service/internal/biz"
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
	return r.data.model.LinkRelation.Create(g.Ident, g.LongUrl)
}

func (r *shortLinkRepo) ShortUrlToLongUrl(ctx context.Context, g *biz.ShortLink) (string, error) {
	url := r.data.model.LinkRelation.GetIdentByLongURL(g.Ident)
	return url, nil
}