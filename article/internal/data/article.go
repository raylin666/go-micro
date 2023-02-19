package data

import (
	"context"
	"fmt"
	"mt/internal/biz"
	"mt/internal/constant/defined"
	"mt/internal/repositories/dbrepo"
	"mt/internal/repositories/dbrepo/model"
	"mt/pkg/logger"
	"mt/pkg/utils"
	"strconv"
)

type articleRepo struct {
	data *Data
	log  *logger.Logger
}

func NewArticleRepo(data *Data, logger *logger.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  logger,
	}
}

// List 文章列表
func (repo *articleRepo) List(ctx context.Context, page, size int32) (m []*model.Article, count int64, err error) {
	var q = dbrepo.NewDefaultDbQuery(repo.data.DbRepo).Article
	var builder = q.WithContext(ctx).Where(q.Status.Eq(model.ENUM_ARTICLE_STATUS_OPEN))
	count, err = builder.Count()
	if err != nil {
		return nil, 0, defined.ErrorDataSelectError
	}

	var offset = 0
	var limit = int(size)
	if int(page) > 1 {
		offset = (int(page) - 1) * limit
	}
	m, err = builder.Limit(int(size)).Offset(offset).Order(q.RecommendFlag.Desc(), q.Sort.Desc(), q.ID.Desc()).Find()
	return m, count, err
}

// CategoryList 文章分类列表
func (repo *articleRepo) CategoryList(ctx context.Context, status int32) (m []*model.ArticleCategory, count int64, err error) {
	var q = dbrepo.NewDefaultDbQuery(repo.data.DbRepo).ArticleCategory
	var builder = q.WithContext(ctx).Where(q.Pid.Eq(0))
	switch status {
	case model.ENUM_ARTICLE_CATEGORY_STATUS_OPEN:
		builder.Where(q.Status.Eq(model.ENUM_ARTICLE_CATEGORY_STATUS_OPEN))
	case model.ENUM_ARTICLE_CATEGORY_STATUS_CLOSE:
		builder.Where(q.Status.Eq(model.ENUM_ARTICLE_CATEGORY_STATUS_CLOSE))
	default:
	}
	count, err = builder.Count()
	if err != nil {
		return nil, 0, defined.ErrorDataSelectError
	}

	m, err = builder.Order(q.Sort.Desc()).Find()
	return m, count, err
}

// ArticleByCategory 获取文章对应分类
func (repo *articleRepo) ArticleByCategory(ctx context.Context, aid int) ([]*model.ArticleCategory, error) {
	if aid <= 0 {
		return nil, defined.ErrorDataNotFound
	}

	var m []*model.ArticleCategory
	var q = dbrepo.NewDefaultDbQuery(repo.data.DbRepo).ArticleCategoryRelation
	var categoryQuery = dbrepo.NewDefaultDbQuery(repo.data.DbRepo).ArticleCategory
	var err = q.WithContext(ctx).Where(q.ArticleId.Eq(aid)).Join(
		categoryQuery,
		categoryQuery.ID.EqCol(q.CategoryId),
		categoryQuery.Status.Eq(model.ENUM_ARTICLE_CATEGORY_STATUS_OPEN),
		categoryQuery.Pid.Eq(0)).Select(
		categoryQuery.ID,
		categoryQuery.Name,
		categoryQuery.Pid,
		categoryQuery.Sort,
		categoryQuery.Color,
	).Scan(&m)

	return m, err
}

// Info 文章详情
func (repo *articleRepo) Info(ctx context.Context, id int) (m *model.Article, em *model.ArticleExtend, err error) {
	if id <= 0 {
		return nil, nil, defined.ErrorDataNotFound
	}

	var q = dbrepo.NewDefaultDbQuery(repo.data.DbRepo).Article
	var extendQuery = dbrepo.NewDefaultDbQuery(repo.data.DbRepo).ArticleExtend
	m, err = q.WithContext(ctx).Unscoped().Where(q.ID.Eq(id)).First()
	if err != nil {
		return nil, nil, defined.ErrorDataSelectError
	}
	em, err = extendQuery.WithContext(ctx).Where(extendQuery.ArticleId.Eq(id)).First()
	if err != nil {
		return nil, nil, defined.ErrorDataSelectError
	}

	return m, em, nil
}

// GetPrevArticle 获取上一篇文章
func (repo *articleRepo) GetPrevArticle(ctx context.Context, id int) *model.Article {
	// 获取ID当前位置
	var sql = fmt.Sprintf("SELECT * FROM (SELECT id, RANK() OVER (ORDER BY recommend_flag DESC, sort DESC, id DESC) AS pos FROM api_article WHERE `status` = %d GROUP BY id ORDER BY recommend_flag DESC, sort DESC, id DESC) ranked WHERE id = %d;", model.ENUM_ARTICLE_STATUS_OPEN, id)
	posResult, err := dbrepo.NewDefaultDb(repo.data.DbRepo).WithContext(ctx).Raw(sql).Rows()
	if err != nil {
		return nil
	}

	var pos int
	var posResultScanRowsMap = utils.ScanRowsMap(posResult)
	for _, posItem := range posResultScanRowsMap {
		pos, _ = strconv.Atoi(posItem["pos"])
		continue
	}

	pos = pos - 2
	if pos < 0 {
		return nil
	}

	var q = dbrepo.NewDefaultDbQuery(repo.data.DbRepo).Article
	m, err := q.WithContext(ctx).Where(q.Status.Eq(model.ENUM_ARTICLE_STATUS_OPEN)).Limit(1).Offset(pos).Order(q.RecommendFlag.Desc(), q.Sort.Desc(), q.ID.Desc()).Select(q.ID, q.Title).First()
	return m
}

// GetNextArticle 获取下一篇文章
func (repo *articleRepo) GetNextArticle(ctx context.Context, id int) *model.Article {
	// 获取ID当前位置
	var sql = fmt.Sprintf("SELECT * FROM (SELECT id, RANK() OVER (ORDER BY recommend_flag DESC, sort DESC, id DESC) AS pos FROM api_article WHERE `status` = %d GROUP BY id ORDER BY recommend_flag DESC, sort DESC, id DESC) ranked WHERE id = %d;", model.ENUM_ARTICLE_STATUS_OPEN, id)
	posResult, err := dbrepo.NewDefaultDb(repo.data.DbRepo).WithContext(ctx).Raw(sql).Rows()
	if err != nil {
		return nil
	}

	var pos int
	var posResultScanRowsMap = utils.ScanRowsMap(posResult)
	for _, posItem := range posResultScanRowsMap {
		pos, _ = strconv.Atoi(posItem["pos"])
		continue
	}

	if pos == 0 {
		return nil
	}

	var q = dbrepo.NewDefaultDbQuery(repo.data.DbRepo).Article
	m, err := q.WithContext(ctx).Where(q.Status.Eq(model.ENUM_ARTICLE_STATUS_OPEN)).Limit(1).Offset(pos).Order(q.RecommendFlag.Desc(), q.Sort.Desc(), q.ID.Desc()).Select(q.ID, q.Title).First()
	return m
}