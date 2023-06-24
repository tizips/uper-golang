package config

import (
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("auth", map[string]any{
		"casbin": map[string]any{
			"table": cfg.Env("auth.casbin.table", "sys_casbin"),
		},
	})
}
