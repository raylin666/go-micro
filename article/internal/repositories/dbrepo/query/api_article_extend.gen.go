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

func newArticleExtend(db *gorm.DB, opts ...gen.DOOption) articleExtend {
	_articleExtend := articleExtend{}

	_articleExtend.articleExtendDo.UseDB(db, opts...)
	_articleExtend.articleExtendDo.UseModel(&model.ArticleExtend{})

	tableName := _articleExtend.articleExtendDo.TableName()
	_articleExtend.ALL = field.NewAsterisk(tableName)
	_articleExtend.ID = field.NewInt(tableName, "id")
	_articleExtend.ArticleId = field.NewInt(tableName, "article_id")
	_articleExtend.Source = field.NewString(tableName, "source")
	_articleExtend.SourceUrl = field.NewString(tableName, "source_url")
	_articleExtend.Content = field.NewString(tableName, "content")
	_articleExtend.Keyword = field.NewString(tableName, "keyword")
	_articleExtend.AttachmentPath = field.NewString(tableName, "attachment_path")

	_articleExtend.fillFieldMap()

	return _articleExtend
}

type articleExtend struct {
	articleExtendDo articleExtendDo

	ALL            field.Asterisk
	ID             field.Int
	ArticleId      field.Int
	Source         field.String
	SourceUrl      field.String
	Content        field.String
	Keyword        field.String
	AttachmentPath field.String

	fieldMap map[string]field.Expr
}

func (a articleExtend) Table(newTableName string) *articleExtend {
	a.articleExtendDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleExtend) As(alias string) *articleExtend {
	a.articleExtendDo.DO = *(a.articleExtendDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleExtend) updateTableName(table string) *articleExtend {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt(table, "id")
	a.ArticleId = field.NewInt(table, "article_id")
	a.Source = field.NewString(table, "source")
	a.SourceUrl = field.NewString(table, "source_url")
	a.Content = field.NewString(table, "content")
	a.Keyword = field.NewString(table, "keyword")
	a.AttachmentPath = field.NewString(table, "attachment_path")

	a.fillFieldMap()

	return a
}

func (a *articleExtend) WithContext(ctx context.Context) *articleExtendDo {
	return a.articleExtendDo.WithContext(ctx)
}

func (a articleExtend) TableName() string { return a.articleExtendDo.TableName() }

func (a articleExtend) Alias() string { return a.articleExtendDo.Alias() }

func (a *articleExtend) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleExtend) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 7)
	a.fieldMap["id"] = a.ID
	a.fieldMap["article_id"] = a.ArticleId
	a.fieldMap["source"] = a.Source
	a.fieldMap["source_url"] = a.SourceUrl
	a.fieldMap["content"] = a.Content
	a.fieldMap["keyword"] = a.Keyword
	a.fieldMap["attachment_path"] = a.AttachmentPath
}

func (a articleExtend) clone(db *gorm.DB) articleExtend {
	a.articleExtendDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleExtend) replaceDB(db *gorm.DB) articleExtend {
	a.articleExtendDo.ReplaceDB(db)
	return a
}

type articleExtendDo struct{ gen.DO }

func (a articleExtendDo) Debug() *articleExtendDo {
	return a.withDO(a.DO.Debug())
}

func (a articleExtendDo) WithContext(ctx context.Context) *articleExtendDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleExtendDo) ReadDB() *articleExtendDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleExtendDo) WriteDB() *articleExtendDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleExtendDo) Session(config *gorm.Session) *articleExtendDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleExtendDo) Clauses(conds ...clause.Expression) *articleExtendDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleExtendDo) Returning(value interface{}, columns ...string) *articleExtendDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleExtendDo) Not(conds ...gen.Condition) *articleExtendDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleExtendDo) Or(conds ...gen.Condition) *articleExtendDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleExtendDo) Select(conds ...field.Expr) *articleExtendDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleExtendDo) Where(conds ...gen.Condition) *articleExtendDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleExtendDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *articleExtendDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a articleExtendDo) Order(conds ...field.Expr) *articleExtendDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleExtendDo) Distinct(cols ...field.Expr) *articleExtendDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleExtendDo) Omit(cols ...field.Expr) *articleExtendDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleExtendDo) Join(table schema.Tabler, on ...field.Expr) *articleExtendDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleExtendDo) LeftJoin(table schema.Tabler, on ...field.Expr) *articleExtendDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleExtendDo) RightJoin(table schema.Tabler, on ...field.Expr) *articleExtendDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleExtendDo) Group(cols ...field.Expr) *articleExtendDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleExtendDo) Having(conds ...gen.Condition) *articleExtendDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleExtendDo) Limit(limit int) *articleExtendDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleExtendDo) Offset(offset int) *articleExtendDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleExtendDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *articleExtendDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleExtendDo) Unscoped() *articleExtendDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleExtendDo) Create(values ...*model.ArticleExtend) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleExtendDo) CreateInBatches(values []*model.ArticleExtend, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleExtendDo) Save(values ...*model.ArticleExtend) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleExtendDo) First() (*model.ArticleExtend, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleExtend), nil
	}
}

func (a articleExtendDo) Take() (*model.ArticleExtend, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleExtend), nil
	}
}

func (a articleExtendDo) Last() (*model.ArticleExtend, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleExtend), nil
	}
}

func (a articleExtendDo) Find() ([]*model.ArticleExtend, error) {
	result, err := a.DO.Find()
	return result.([]*model.ArticleExtend), err
}

func (a articleExtendDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArticleExtend, err error) {
	buf := make([]*model.ArticleExtend, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleExtendDo) FindInBatches(result *[]*model.ArticleExtend, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleExtendDo) Attrs(attrs ...field.AssignExpr) *articleExtendDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleExtendDo) Assign(attrs ...field.AssignExpr) *articleExtendDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleExtendDo) Joins(fields ...field.RelationField) *articleExtendDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleExtendDo) Preload(fields ...field.RelationField) *articleExtendDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleExtendDo) FirstOrInit() (*model.ArticleExtend, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleExtend), nil
	}
}

func (a articleExtendDo) FirstOrCreate() (*model.ArticleExtend, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleExtend), nil
	}
}

func (a articleExtendDo) FindByPage(offset int, limit int) (result []*model.ArticleExtend, count int64, err error) {
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

func (a articleExtendDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleExtendDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleExtendDo) Delete(models ...*model.ArticleExtend) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleExtendDo) withDO(do gen.Dao) *articleExtendDo {
	a.DO = *do.(*gen.DO)
	return a
}
