package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableBlgArticle = "blg_article"

type BlgArticle struct {
	ID         string          `gorm:"column:id;primaryKey"`
	CategoryID string          `gorm:"column:category_id"`
	UserID     string          `gorm:"column:user_id"`
	Name       string          `gorm:"column:name"`
	Picture    string          `gorm:"column:picture"`
	Source     string          `gorm:"column:source"`
	URL        string          `gorm:"column:url"`
	Summary    string          `gorm:"column:summary"`
	IsComment  uint8           `gorm:"column:is_comment"`
	IsEnable   uint8           `gorm:"column:is_enable"`
	CreatedAt  carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt  carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt  `gorm:"column:deleted_at"`

	Category *BlgCategory `gorm:"foreignKey:ID;references:CategoryID"`
	SEO      *BlgSEO      `gorm:"foreignKey:OtherID;references:ID"`
	HTML     *BlgHTML     `gorm:"foreignKey:OtherID;references:ID"`
}

func (BlgArticle) TableName() string {
	return TableBlgArticle
}
