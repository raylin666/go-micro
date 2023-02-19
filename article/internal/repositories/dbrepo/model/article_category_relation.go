package model

type ArticleCategoryRelation struct {
	ID         int `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键ID
	ArticleId  int `gorm:"column:article_id" json:"article_id"`               // 文章ID
	CategoryId int `gorm:"column:category_id" json:"category_id"`             // 分类ID
}
