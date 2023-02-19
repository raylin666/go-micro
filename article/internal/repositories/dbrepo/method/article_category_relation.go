package method

import "gorm.io/gen"

type ArticleCategoryRelation interface {
	// where("`article_id`=@article_id and `category_id`=@category_id")
	FirstArticleAndCategory(article_id, category_id string) (gen.T, error)
}
