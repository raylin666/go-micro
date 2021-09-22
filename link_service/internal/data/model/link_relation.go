package model

import (
	"gorm.io/gorm"
	"time"
)

type LinkRelation struct {
	ID        int64     `json:"id"`
	Ident     string    `json:"ident"`
	LongUrl   string    `json:"long_url"`
	State     int8      `json:"state"`
	CreatedAt time.Time `json:"created_at"`
}

func (LinkRelation) TableName() string {
	return "link_relation"
}

func (m LinkRelation) Create(db *gorm.DB, ident string, longUrl string) error {
	result := db.Model(m).Create(&LinkRelation{
		Ident:     ident,
		LongUrl:   longUrl,
		State:     1,
		CreatedAt: time.Now(),
	})

	return result.Error
}
