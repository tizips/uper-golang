package site

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	authConstants "github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/uper-go/admin/constants"
	"github.com/tizips/uper-go/admin/helper/authorize"
	req "github.com/tizips/uper-go/admin/http/request/site"
	res "github.com/tizips/uper-go/admin/http/response/site"
	"github.com/tizips/uper-go/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ToRoleByPaginate(ctx context.Context, c *app.RequestContext) {

	var request req.ToRoleByPaginate

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	responses := response.Paginate[res.ToRoleByPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm

	tx.Model(&model.SysRole{}).Count(&responses.Total)

	if responses.Total > 0 {

		var roles []model.SysRole

		tx.
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Order("`id` desc").
			Find(&roles)

		responses.Data = make([]res.ToRoleByPaginate, len(roles))

		for index, item := range roles {
			responses.Data[index] = res.ToRoleByPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Summary:   item.Summary,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(c, responses)
}

func ToRoleByInformation(ctx context.Context, c *app.RequestContext) {

	id := c.Param("id")

	var role model.SysRole

	fr := facades.Gorm.Preload("BindPermissions").First(&role, "`id`=?", id)
	if errors.Is(fr.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该数据")
		return
	} else if fr.Error != nil {
		http.Fail(c, "查询失败：%v", fr.Error)
		return
	}

	responses := res.ToRoleByInformation{
		ID:          role.ID,
		Name:        role.Name,
		Permissions: make([]string, len(role.BindPermissions)),
		Summary:     role.Summary,
		CreatedAt:   role.CreatedAt.ToDateTimeString(),
	}

	for index, item := range role.BindPermissions {
		responses.Permissions[index] = item.Permission
	}

	http.Success(c, responses)
}

func DoRoleByCreate(ctx context.Context, c *app.RequestContext) {

	var request req.DoRoleByCreate

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	results := auth.HandlerPermissionsByParent(constants.Permissions, 0, request.Permissions)

	tx := facades.Gorm.Begin()

	role := model.SysRole{
		Name:    request.Name,
		Summary: request.Summary,
	}

	if cr := tx.Create(&role); cr.Error != nil {
		tx.Rollback()
		http.Fail(c, "创建失败：%v", cr.Error)
		return
	}

	bindings := make([]model.SysRoleBindPermission, len(results))
	permissions := make([][]string, len(results))

	for index, item := range results {
		bindings[index] = model.SysRoleBindPermission{
			RoleID:     role.ID,
			Permission: item.Code,
		}
		permissions[index] = []string{item.Code}
	}

	if cb := tx.CreateInBatches(&bindings, 50); cb.Error != nil {
		tx.Rollback()
		http.Fail(c, "创建失败：%v", cb.Error)
		return
	}

	if ok, _ := facades.Casbin.AddPermissionsForUser(auth.NameOfRole(role.ID), permissions...); !ok {
		tx.Rollback()
		http.Fail(c, "创建失败")
		return
	}

	tx.Commit()

	http.Success[any](c)
}

func DoRoleByUpdate(ctx context.Context, c *app.RequestContext) {

	var request req.DoRoleByUpdate

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	var role model.SysRole

	fr := facades.Gorm.Preload("BindPermissions").First(&role, "`id`=?", request.ID)
	if errors.Is(fr.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该数据")
		return
	} else if fr.Error != nil {
		http.Fail(c, "查询失败：%v", fr.Error)
		return
	}

	if role.ID == authConstants.CodeOfDeveloper {
		http.Fail(c, "内置角色，无法修改")
		return
	}

	//	查询出真实可用的权限
	results := auth.HandlerPermissionsByParent(constants.Permissions, 0, request.Permissions)

	var creates []model.SysRoleBindPermission
	var deletes []uint

	for _, item := range results {

		mark := true

		for _, value := range role.BindPermissions {
			if item.Code == value.Permission {
				mark = false
				break
			}
		}

		if mark {
			creates = append(creates, model.SysRoleBindPermission{
				RoleID:     role.ID,
				Permission: item.Code,
			})
		}
	}

	for _, item := range role.BindPermissions {

		mark := true

		for _, value := range results {
			if item.Permission == value.Code {
				mark = false
				break
			}
		}

		if mark {
			deletes = append(deletes, item.ID)
		}
	}

	tx := facades.Gorm.Begin()

	role.Name = request.Name
	role.Summary = request.Summary

	if ur := tx.Omit(clause.Associations).Save(&role); ur.Error != nil {
		tx.Rollback()
		http.Fail(c, "修改失败：%v", ur.Error)
		return
	}

	if len(creates) > 0 {
		if cbp := tx.CreateInBatches(&creates, 50); cbp.Error != nil {
			tx.Rollback()
			http.Fail(c, "修改失败：%v", cbp.Error)
			return
		}
	}

	if len(deletes) > 0 {
		if cbp := tx.Delete(&model.SysRoleBindPermission{}, "`role_id`=? and `id` IN (?)", role.ID, deletes); cbp.Error != nil {
			tx.Rollback()
			http.Fail(c, "修改失败：%v", cbp.Error)
			return
		}
	}

	if len(creates) > 0 || len(deletes) > 0 {

		//	清除旧权限
		if ok, _ := facades.Casbin.DeletePermissionsForUser(auth.NameOfRole(role.ID)); !ok {
			tx.Rollback()
			http.Fail(c, "修改失败")
			return
		}

		//	生成新权限
		var permissions []model.SysRoleBindPermission

		tx.Find(&permissions, "`role_id`=?", role.ID)

		if len(permissions) > 0 {

			items := make([][]string, len(permissions))

			for index, item := range permissions {
				items[index] = []string{item.Permission}
			}

			if ok, _ := facades.Casbin.AddPermissionsForUser(auth.NameOfRole(role.ID), items...); !ok {
				tx.Rollback()
				http.Fail(c, "修改失败")
				return
			}
		}
	}

	tx.Commit()

	http.Success[any](c)
}

func DoRoleByDelete(ctx context.Context, c *app.RequestContext) {

	id := c.Param("id")

	var role model.SysRole

	fr := facades.Gorm.Preload("BindPermissions").First(&role, "`id`=?", id)
	if errors.Is(fr.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该数据")
		return
	} else if fr.Error != nil {
		http.Fail(c, "查询失败：%v", fr.Error)
		return
	}

	if role.ID == authConstants.CodeOfDeveloper {
		http.Fail(c, "内置角色，无法删除")
		return
	}

	tx := facades.Gorm.Begin()

	if dr := tx.Delete(&role, "`id`=?", role.ID); dr.Error != nil {
		tx.Rollback()
		http.Fail(c, "删除失败：%v", dr.Error)
		return
	}

	//	清除旧角色
	if ok, _ := facades.Casbin.DeleteRole(auth.NameOfRole(role.ID)); !ok {
		tx.Rollback()
		http.Fail(c, "删除失败")
		return
	}

	tx.Commit()

	http.Success[any](c)
}

func ToRoleByOpening(ctx context.Context, c *app.RequestContext) {

	var roles []model.SysRole

	tx := facades.Gorm

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(authorize.ID(c)), auth.NameOfRoleWithDeveloper()); !ok {
		tx = tx.Where("`id`<>?", authConstants.CodeOfDeveloper)
	}

	tx.Find(&roles)

	responses := make([]response.Opening[uint], len(roles))

	for index, item := range roles {
		responses[index] = response.Opening[uint]{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	http.Success(c, responses)
}
