package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableSysRole = "sys_role"

type SysRole struct {
	ID        uint            `gorm:"column:id;primaryKey"`
	Name      string          `gorm:"column:name"`
	Summary   string          `gorm:"column:summary"`
	CreatedAt carbon.DateTime `gorm:"column:created_at"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`

	BindPermissions []SysRoleBindPermission `gorm:"foreignKey:RoleID;references:ID"`
}

func (SysRole) TableName() string {
	return TableSysRole
}
