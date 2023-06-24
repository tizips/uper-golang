package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableBlgCategory = "blg_category"

type BlgCategory struct {
	ID        string          `gorm:"column:id;primaryKey"`
	ParentID  *string         `gorm:"column:parent_id"`
	Type      string          `gorm:"column:type"`
	Name      string          `gorm:"column:name"`
	Picture   string          `gorm:"column:picture"`
	Order     uint8           `gorm:"column:order"`
	IsComment uint8           `gorm:"column:is_comment"`
	IsEnable  uint8           `gorm:"column:is_enable"`
	CreatedAt carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`

	SEO  *BlgSEO  `gorm:"foreignKey:OtherID;references:ID"`
	HTML *BlgHTML `gorm:"foreignKey:OtherID;references:ID"`
}

func (BlgCategory) TableName() string {
	return TableBlgCategory
}

const (
	BlogCategoryForTypeOfParent = "parent"
	BlogCategoryForTypeOfPage   = "page"
	BlogCategoryForTypeOfList   = "list"
)
