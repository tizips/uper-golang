package config

import (
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/console"
	"github.com/herhe-com/framework/console/consoles"
	cons "github.com/herhe-com/framework/contracts/console"
	"github.com/herhe-com/framework/contracts/service"
	"github.com/herhe-com/framework/database/gorm"
	"github.com/herhe-com/framework/database/redis"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/filesystem"
	"github.com/herhe-com/framework/microservice/locker"
	"github.com/herhe-com/framework/microservice/snowflake"
	"github.com/herhe-com/framework/validation"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("app", map[string]any{
		"name":     cfg.Env("app.name", "UPER"),
		"address":  cfg.Env("app.address", "0.0.0.0"),
		"port":     cfg.Env("app.port", "9600"),
		"node":     cfg.Env("app.node", 1),
		"debug":    cfg.Env("app.mode", true),
		"domain":   cfg.Env("app.domain", "http://127.0.0.1:9600"),
		"location": cfg.Env("app.location", "Asia/Shanghai"),
		"providers": []service.Provider{
			&gorm.ServiceProvider{},
			&redis.ServiceProvider{},
			&filesystem.ServiceProvider{},
			&snowflake.ServiceProvider{},
			&locker.ServiceProvider{},
			&validation.ServiceProvider{},
			&auth.ServiceProvider{},
			&console.ServiceProvider{},
		},
		"consoles": []cons.Provider{
			&consoles.MigrationProvider{},
			&consoles.ServerProvider{},
			//&consoles2.RoleProvider{},
			//&consoles2.DeveloperProvider{},
		},
	})
}

func Boot() {

}
