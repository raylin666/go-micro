// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                       = new(Query)
	Article                 *article
	ArticleCategory         *articleCategory
	ArticleCategoryRelation *articleCategoryRelation
	ArticleExtend           *articleExtend
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Article = &Q.Article
	ArticleCategory = &Q.ArticleCategory
	ArticleCategoryRelation = &Q.ArticleCategoryRelation
	ArticleExtend = &Q.ArticleExtend
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                      db,
		Article:                 newArticle(db, opts...),
		ArticleCategory:         newArticleCategory(db, opts...),
		ArticleCategoryRelation: newArticleCategoryRelation(db, opts...),
		ArticleExtend:           newArticleExtend(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Article                 article
	ArticleCategory         articleCategory
	ArticleCategoryRelation articleCategoryRelation
	ArticleExtend           articleExtend
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                      db,
		Article:                 q.Article.clone(db),
		ArticleCategory:         q.ArticleCategory.clone(db),
		ArticleCategoryRelation: q.ArticleCategoryRelation.clone(db),
		ArticleExtend:           q.ArticleExtend.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                      db,
		Article:                 q.Article.replaceDB(db),
		ArticleCategory:         q.ArticleCategory.replaceDB(db),
		ArticleCategoryRelation: q.ArticleCategoryRelation.replaceDB(db),
		ArticleExtend:           q.ArticleExtend.replaceDB(db),
	}
}

type queryCtx struct {
	Article                 *articleDo
	ArticleCategory         *articleCategoryDo
	ArticleCategoryRelation *articleCategoryRelationDo
	ArticleExtend           *articleExtendDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Article:                 q.Article.WithContext(ctx),
		ArticleCategory:         q.ArticleCategory.WithContext(ctx),
		ArticleCategoryRelation: q.ArticleCategoryRelation.WithContext(ctx),
		ArticleExtend:           q.ArticleExtend.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
