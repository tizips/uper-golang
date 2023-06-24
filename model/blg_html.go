package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableBlgHTML = "blg_html"

type BlgHTML struct {
	ID        uint            `gorm:"column:id;primaryKey"`
	Type      string          `gorm:"column:type"`
	OtherID   string          `gorm:"column:other_id"`
	Content   string          `gorm:"column:content"`
	Text      string          `gorm:"column:text"`
	CreatedAt carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (BlgHTML) TableName() string {
	return TableBlgHTML
}

const (
	BlogArticleForSearchIndex = "article"
)
