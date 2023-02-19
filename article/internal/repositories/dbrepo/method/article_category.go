package method

import (
	"gorm.io/gen"
)

type ArticleCategory interface {
	// where("`name`=@name")
	FindByName(name string) (gen.T, error)
}
