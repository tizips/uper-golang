package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableSysRoleBindPermission = "sys_role_bind_permission"

type SysRoleBindPermission struct {
	ID         uint            `gorm:"column:id;primaryKey"`
	RoleID     uint            `gorm:"column:role_id"`
	Permission string          `gorm:"column:permission"`
	CreatedAt  carbon.DateTime `gorm:"column:created_at"`
	DeletedAt  gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (SysRoleBindPermission) TableName() string {
	return TableSysRoleBindPermission
}
