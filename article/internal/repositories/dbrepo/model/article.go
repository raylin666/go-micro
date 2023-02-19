package model

const (
	// ENUM_ARTICLE_STATUS_CLOSE 状态关闭
	ENUM_ARTICLE_STATUS_CLOSE = 0
	// ENUM_ARTICLE_STATUS_OPEN 状态开启
	ENUM_ARTICLE_STATUS_OPEN = 1
)

type Article struct {
	Title           string `gorm:"column:title" json:"title"`                         // 文章标题
	Author          string `gorm:"column:author" json:"author"`                       // 文章作者
	Summary         string `gorm:"column:summary" json:"summary"`                     // 文章摘要
	Cover           string `gorm:"column:cover" json:"cover"`                         // 文章封面图片
	Sort            int32  `gorm:"column:sort" json:"sort"`                           // 排序
	RecommendFlag   int8   `gorm:"column:recommend_flag" json:"recommend_flag"`       // 文章推荐标识 0:未推荐，1:已推荐
	CommentedFlag   int8   `gorm:"column:commented_flag" json:"commented_flag"`       // 文章是否允许评论 1:允许，0:不允许
	Status          int8   `gorm:"column:status" json:"status"`                       // 状态 0:已关闭 1:已开启
	ViewCount       int    `gorm:"column:view_count" json:"view_count"`               // 文章浏览量
	CommentCount    int    `gorm:"column:comment_count" json:"comment_count"`         // 文章评论数
	CollectionCount int    `gorm:"column:collection_count" json:"collection_count"`	  // 文章收藏量
	ZanCount 	    int    `gorm:"column:zan_count" json:"zan_count"`	  			  // 文章点赞数
	ShareCount      int    `gorm:"column:share_count" json:"share_count"`             // 文章分享数
	UserId          int    `gorm:"column:user_id" json:"user_id"`                     // 发布者编号
	LastCommentedAt int64  `gorm:"column:last_commented_at" json:"last_commented_at"` // 最新评论时间

	BaseModel
}
