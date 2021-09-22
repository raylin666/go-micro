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

type LinkRelationModel struct {
	db    *gorm.DB
	model *LinkRelation
}

func (*LinkRelation) TableName() string {
	return "link_relation"
}

func NewLinkRelationModel(db *gorm.DB) *LinkRelationModel {
	return &LinkRelationModel{db: db}
}

func (m *LinkRelationModel) Create(ident string, longUrl string) error {
	result := m.db.Model(m.model).Create(&LinkRelation{
		Ident:     ident,
		LongUrl:   longUrl,
		State:     1,
		CreatedAt: time.Now(),
	})

	return result.Error
}

func (m *LinkRelationModel) GetIdentByLongURL(ident string) string {
	var dest LinkRelation
	m.db.Model(m.model).Where("ident = ?", ident).Find(&dest)
	return dest.LongUrl
}
