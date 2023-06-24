package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableBlgSEO = "blg_seo"

type BlgSEO struct {
	ID          uint            `gorm:"column:id;primaryKey"`
	Type        string          `gorm:"column:type"`
	OtherID     string          `gorm:"column:other_id"`
	Title       string          `gorm:"column:title"`
	Keyword     string          `gorm:"column:keyword"`
	Description string          `gorm:"column:description"`
	CreatedAt   carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt   carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (BlgSEO) TableName() string {
	return TableBlgSEO
}
