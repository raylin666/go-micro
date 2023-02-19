package model

const (
	// ENUM_ARTICLE_CATEGORY_STATUS_CLOSE 状态关闭
	ENUM_ARTICLE_CATEGORY_STATUS_CLOSE = 0
	// ENUM_ARTICLE_CATEGORY_STATUS_OPEN 状态开启
	ENUM_ARTICLE_CATEGORY_STATUS_OPEN = 1
)

type ArticleCategory struct {
	Pid    int    `gorm:"column:pid" json:"pid"`       // 上级分类
	Name   string `gorm:"column:name" json:"name"`     // 分类名称
	Color  string `gorm:"column:color" json:"color"`   // 分类颜色
	Sort   int    `gorm:"column:sort" json:"sort"`     // 排序
	Status int8   `gorm:"column:status" json:"status"` // 状态 0:已关闭 1:已开启

	BaseModel
}
