package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableBlgLink = "blg_link"

type BlgLink struct {
	ID        uint            `gorm:"column:id;primaryKey"`
	Name      string          `gorm:"column:name"`
	URL       string          `gorm:"column:url"`
	Logo      string          `gorm:"column:logo"`
	Email     string          `gorm:"column:email"`
	Position  string          `gorm:"column:position"`
	Order     uint8           `gorm:"column:order"`
	IsEnable  uint8           `gorm:"column:is_enable"`
	CreatedAt carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (BlgLink) TableName() string {
	return TableBlgLink
}

const (
	BlgLinkForPositionOfAll    = "all"
	BlgLinkForPositionOfBottom = "bottom"
	BlgLinkForPositionOfOther  = "other"
)
