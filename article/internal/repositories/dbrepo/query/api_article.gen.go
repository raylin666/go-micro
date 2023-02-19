// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"mt/internal/repositories/dbrepo/model"
)

func newArticle(db *gorm.DB, opts ...gen.DOOption) article {
	_article := article{}

	_article.articleDo.UseDB(db, opts...)
	_article.articleDo.UseModel(&model.Article{})

	tableName := _article.articleDo.TableName()
	_article.ALL = field.NewAsterisk(tableName)
	_article.Title = field.NewString(tableName, "title")
	_article.Author = field.NewString(tableName, "author")
	_article.Summary = field.NewString(tableName, "summary")
	_article.Cover = field.NewString(tableName, "cover")
	_article.Sort = field.NewInt32(tableName, "sort")
	_article.RecommendFlag = field.NewInt8(tableName, "recommend_flag")
	_article.CommentedFlag = field.NewInt8(tableName, "commented_flag")
	_article.Status = field.NewInt8(tableName, "status")
	_article.ViewCount = field.NewInt(tableName, "view_count")
	_article.CommentCount = field.NewInt(tableName, "comment_count")
	_article.CollectionCount = field.NewInt(tableName, "collection_count")
	_article.ZanCount = field.NewInt(tableName, "zan_count")
	_article.ShareCount = field.NewInt(tableName, "share_count")
	_article.UserId = field.NewInt(tableName, "user_id")
	_article.LastCommentedAt = field.NewInt64(tableName, "last_commented_at")
	_article.ID = field.NewInt(tableName, "id")
	_article.CreatedAt = field.NewInt64(tableName, "created_at")
	_article.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_article.DeletedAt = field.NewField(tableName, "deleted_at")

	_article.fillFieldMap()

	return _article
}

type article struct {
	articleDo articleDo

	ALL             field.Asterisk
	Title           field.String
	Author          field.String
	Summary         field.String
	Cover           field.String
	Sort            field.Int32
	RecommendFlag   field.Int8
	CommentedFlag   field.Int8
	Status          field.Int8
	ViewCount       field.Int
	CommentCount    field.Int
	CollectionCount field.Int
	ZanCount        field.Int
	ShareCount      field.Int
	UserId          field.Int
	LastCommentedAt field.Int64
	ID              field.Int
	CreatedAt       field.Int64
	UpdatedAt       field.Int64
	DeletedAt       field.Field

	fieldMap map[string]field.Expr
}

func (a article) Table(newTableName string) *article {
	a.articleDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a article) As(alias string) *article {
	a.articleDo.DO = *(a.articleDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *article) updateTableName(table string) *article {
	a.ALL = field.NewAsterisk(table)
	a.Title = field.NewString(table, "title")
	a.Author = field.NewString(table, "author")
	a.Summary = field.NewString(table, "summary")
	a.Cover = field.NewString(table, "cover")
	a.Sort = field.NewInt32(table, "sort")
	a.RecommendFlag = field.NewInt8(table, "recommend_flag")
	a.CommentedFlag = field.NewInt8(table, "commented_flag")
	a.Status = field.NewInt8(table, "status")
	a.ViewCount = field.NewInt(table, "view_count")
	a.CommentCount = field.NewInt(table, "comment_count")
	a.CollectionCount = field.NewInt(table, "collection_count")
	a.ZanCount = field.NewInt(table, "zan_count")
	a.ShareCount = field.NewInt(table, "share_count")
	a.UserId = field.NewInt(table, "user_id")
	a.LastCommentedAt = field.NewInt64(table, "last_commented_at")
	a.ID = field.NewInt(table, "id")
	a.CreatedAt = field.NewInt64(table, "created_at")
	a.UpdatedAt = field.NewInt64(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *article) WithContext(ctx context.Context) *articleDo { return a.articleDo.WithContext(ctx) }

func (a article) TableName() string { return a.articleDo.TableName() }

func (a article) Alias() string { return a.articleDo.Alias() }

func (a *article) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *article) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 19)
	a.fieldMap["title"] = a.Title
	a.fieldMap["author"] = a.Author
	a.fieldMap["summary"] = a.Summary
	a.fieldMap["cover"] = a.Cover
	a.fieldMap["sort"] = a.Sort
	a.fieldMap["recommend_flag"] = a.RecommendFlag
	a.fieldMap["commented_flag"] = a.CommentedFlag
	a.fieldMap["status"] = a.Status
	a.fieldMap["view_count"] = a.ViewCount
	a.fieldMap["comment_count"] = a.CommentCount
	a.fieldMap["collection_count"] = a.CollectionCount
	a.fieldMap["zan_count"] = a.ZanCount
	a.fieldMap["share_count"] = a.ShareCount
	a.fieldMap["user_id"] = a.UserId
	a.fieldMap["last_commented_at"] = a.LastCommentedAt
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
}

func (a article) clone(db *gorm.DB) article {
	a.articleDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a article) replaceDB(db *gorm.DB) article {
	a.articleDo.ReplaceDB(db)
	return a
}

type articleDo struct{ gen.DO }

func (a articleDo) Debug() *articleDo {
	return a.withDO(a.DO.Debug())
}

func (a articleDo) WithContext(ctx context.Context) *articleDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleDo) ReadDB() *articleDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleDo) WriteDB() *articleDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleDo) Session(config *gorm.Session) *articleDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleDo) Clauses(conds ...clause.Expression) *articleDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleDo) Returning(value interface{}, columns ...string) *articleDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleDo) Not(conds ...gen.Condition) *articleDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleDo) Or(conds ...gen.Condition) *articleDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleDo) Select(conds ...field.Expr) *articleDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleDo) Where(conds ...gen.Condition) *articleDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *articleDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a articleDo) Order(conds ...field.Expr) *articleDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleDo) Distinct(cols ...field.Expr) *articleDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleDo) Omit(cols ...field.Expr) *articleDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleDo) Join(table schema.Tabler, on ...field.Expr) *articleDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *articleDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleDo) RightJoin(table schema.Tabler, on ...field.Expr) *articleDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleDo) Group(cols ...field.Expr) *articleDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleDo) Having(conds ...gen.Condition) *articleDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleDo) Limit(limit int) *articleDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleDo) Offset(offset int) *articleDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *articleDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleDo) Unscoped() *articleDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleDo) Create(values ...*model.Article) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleDo) CreateInBatches(values []*model.Article, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleDo) Save(values ...*model.Article) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleDo) First() (*model.Article, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) Take() (*model.Article, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) Last() (*model.Article, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) Find() ([]*model.Article, error) {
	result, err := a.DO.Find()
	return result.([]*model.Article), err
}

func (a articleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Article, err error) {
	buf := make([]*model.Article, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleDo) FindInBatches(result *[]*model.Article, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleDo) Attrs(attrs ...field.AssignExpr) *articleDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleDo) Assign(attrs ...field.AssignExpr) *articleDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleDo) Joins(fields ...field.RelationField) *articleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleDo) Preload(fields ...field.RelationField) *articleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleDo) FirstOrInit() (*model.Article, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) FirstOrCreate() (*model.Article, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Article), nil
	}
}

func (a articleDo) FindByPage(offset int, limit int) (result []*model.Article, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a articleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleDo) Delete(models ...*model.Article) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleDo) withDO(do gen.Dao) *articleDo {
	a.DO = *do.(*gen.DO)
	return a
}
