package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键ID
	CreatedAt int64          `gorm:"column:created_at" json:"created_at"`               // 创建时间
	UpdatedAt int64          `gorm:"column:updated_at" json:"updated_at"`               // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`               // 删除时间
}
