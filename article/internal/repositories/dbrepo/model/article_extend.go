package model

type ArticleExtend struct {
	ID             int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键ID
	ArticleId      int    `gorm:"column:article_id" json:"article_id"`               // 文章ID
	Source         string `gorm:"column:source" json:"source"`                       // 文章来源
	SourceUrl      string `gorm:"column:source_url" json:"source_url"`               // 文章来源链接
	Content        string `gorm:"column:content" json:"content"`                     // 文章正文
	Keyword        string `gorm:"column:keyword" json:"keyword"`                     // 文章关键词
	AttachmentPath string `gorm:"column:attachment_path" json:"attachment_path"`     // 文章附件路径
}
