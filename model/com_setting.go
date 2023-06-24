package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableComSetting = "com_setting"

type ComSetting struct {
	ID         uint            `gorm:"column:id;primaryKey"`
	Module     string          `gorm:"column:module"`
	Type       string          `gorm:"column:type"`
	Label      string          `gorm:"column:label"`
	Key        string          `gorm:"column:key"`
	Val        string          `gorm:"column:val"`
	IsRequired uint8           `gorm:"column:is_required"`
	Order      uint8           `gorm:"column:order"`
	CreatedAt  carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt  carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (ComSetting) TableName() string {
	return TableComSetting
}

const (
	ComSettingForTypeOfInput    = "input"
	ComSettingForTypeOfEnable   = "enable"
	ComSettingForTypeOfURL      = "url"
	ComSettingForTypeOfEmail    = "email"
	ComSettingForTypeOfPicture  = "picture"
	ComSettingForTypeOfTextarea = "textarea"

	ComSettingForIsRequiredOfYes = 1
	ComSettingForIsRequiredOfNo  = 2
)
