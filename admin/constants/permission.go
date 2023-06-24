package constants

import "github.com/herhe-com/framework/contracts/auth"

var Permissions = []auth.Permission{
	{
		Code: "site",
		Name: "站点",
		Children: []auth.Permission{
			{
				Code: "role",
				Name: "角色",
				Children: []auth.Permission{
					{
						Code:   "create",
						Name:   "创建",
						Common: true,
					},
					{
						Code:   "update",
						Name:   "修改",
						Common: true,
					},
					{
						Code:   "delete",
						Name:   "删除",
						Common: true,
					},
					{
						Code:   "paginate",
						Name:   "列表",
						Common: true,
					},
				},
			},
			{
				Code: "user",
				Name: "账号",
				Children: []auth.Permission{
					{
						Code:   "create",
						Name:   "创建",
						Common: true,
					},
					{
						Code:   "update",
						Name:   "修改",
						Common: true,
					},
					{
						Code:   "delete",
						Name:   "删除",
						Common: true,
					},
					{
						Code:   "enable",
						Name:   "启禁",
						Common: true,
					},
					{
						Code:   "paginate",
						Name:   "列表",
						Common: true,
					},
				},
			},
		},
	},
	{
		Code: "blog",
		Name: "博客",
		Children: []auth.Permission{
			{
				Code: "category",
				Name: "栏目",
				Children: []auth.Permission{
					{
						Code:   "create",
						Name:   "创建",
						Common: true,
					},
					{
						Code:   "update",
						Name:   "修改",
						Common: true,
					},
					{
						Code:   "delete",
						Name:   "删除",
						Common: true,
					},
					{
						Code:   "enable",
						Name:   "启禁",
						Common: true,
					},
					{
						Code:   "tree",
						Name:   "列表",
						Common: true,
					},
				},
			},
			{
				Code: "article",
				Name: "文章",
				Children: []auth.Permission{
					{
						Code:   "create",
						Name:   "创建",
						Common: true,
					},
					{
						Code:   "update",
						Name:   "修改",
						Common: true,
					},
					{
						Code:   "delete",
						Name:   "删除",
						Common: true,
					},
					{
						Code:   "enable",
						Name:   "启禁",
						Common: true,
					},
					{
						Code:   "paginate",
						Name:   "列表",
						Common: true,
					},
				},
			},
			{
				Code: "link",
				Name: "友链",
				Children: []auth.Permission{
					{
						Code:   "create",
						Name:   "创建",
						Common: true,
					},
					{
						Code:   "update",
						Name:   "修改",
						Common: true,
					},
					{
						Code:   "delete",
						Name:   "删除",
						Common: true,
					},
					{
						Code:   "enable",
						Name:   "启禁",
						Common: true,
					},
					{
						Code:   "paginate",
						Name:   "列表",
						Common: true,
					},
				},
			},
			{
				Code: "setting",
				Name: "设置",
				Children: []auth.Permission{
					{
						Code:   "update",
						Name:   "修改",
						Common: true,
					},
					{
						Code:   "list",
						Name:   "列表",
						Common: true,
					},
				},
			},
		},
	},
}
