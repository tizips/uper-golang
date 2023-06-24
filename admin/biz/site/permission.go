package site

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	authConstants "github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/uper-go/admin/constants"
	"github.com/tizips/uper-go/admin/helper/authorize"
	"github.com/tizips/uper-go/model"
)

func ToPermissions(ctx context.Context, c *app.RequestContext) {

	var results []authConstants.PermissionOfTrees

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(authorize.ID(c)), auth.NameOfRoleWithDeveloper()); ok {
		results = auth.HandlerPermissionsByTrees(constants.Permissions, 0, nil, nil, true)
	} else {

		var bindings []model.SysRoleBindPermission

		facades.Gorm.
			Where("exists (?)", facades.Gorm.
				Select("1").
				Model(&model.SysUserBindRole{}).
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), authorize.ID(c)).
				Where("exists (?)", facades.Gorm.
					Select("1").
					Model(&model.SysUser{}).
					Where(fmt.Sprintf("%s.user_id=%s.id  and `is_enable`=?", model.TableSysUserBindRole, model.TableSysUser), util.EnableOfYes),
				),
			).
			Find(&bindings)

		codes := make([]string, len(bindings))

		for index, item := range bindings {
			codes[index] = item.Permission
		}

		results = auth.HandlerPermissionsByTrees(constants.Permissions, 0, nil, codes, false)
	}

	http.Success(c, results)
}
