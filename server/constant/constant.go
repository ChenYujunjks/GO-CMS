package constant

import "github.com/Xi-Yuer/cms/config"

var AUTHORIZATION = `Authorization`

var JWTPAYLOAD = `JwtPayload`

var PermissionMap = map[string]string{
	// 用户管理
	"^GET:" + config.Config.APP.BASEURL + "/users$":            `GET:/users`,
	"^POST:" + config.Config.APP.BASEURL + "/users$":           `POST:/users`,
	"^GET:" + config.Config.APP.BASEURL + "/users/roles/\\d+$": `GET:/users/roles/:id`,
	"^GET:" + config.Config.APP.BASEURL + "/users/\\d+$":       `GET:/users/:id`,
	"^PATCH:" + config.Config.APP.BASEURL + "/users/\\d+$":     `PATCH:/users/:id`,
	"^DELETE:" + config.Config.APP.BASEURL + "/users/\\d+$":    `DELETE:/users/:id`,
	"^POST:" + config.Config.APP.BASEURL + "/users/export$":    `POST:/users/export`,

	// 角色管理
	"^GET:" + config.Config.APP.BASEURL + "/roles$":           `GET:/roles`,
	"^GET:" + config.Config.APP.BASEURL + "/users/role/\\d+$": `GET:/users/role/:id`,
	"^POST:" + config.Config.APP.BASEURL + "/roles$":          `POST:/roles`,
	"^PATCH:" + config.Config.APP.BASEURL + "/roles/\\d+$":    `PATCH:/roles/:id`,
	"^DELETE:" + config.Config.APP.BASEURL + "/roles/\\d+$":   `DELETE:/roles/:id`,
	"^POST:" + config.Config.APP.BASEURL + "/roles/export$":   `POST:/roles/export`,

	// 部门管理
	"^GET:" + config.Config.APP.BASEURL + "/department$":         `GET:/department`,
	"^POST:" + config.Config.APP.BASEURL + "/department$":        `POST:/department`,
	"^PATCH:" + config.Config.APP.BASEURL + "/department/\\d+$":  `PATCH:/department/:id`,
	"^DELETE:" + config.Config.APP.BASEURL + "/department/\\d+$": `DELETE:/department/:id`,

	// 菜单管理
	"^GET:" + config.Config.APP.BASEURL + "/pages$":         `GET:/pages`,
	"^POST:" + config.Config.APP.BASEURL + "/pages$":        `POST:/pages`,
	"^PATCH:" + config.Config.APP.BASEURL + "/pages/\\d+$":  `PATCH:/pages/:id`,
	"^DELETE:" + config.Config.APP.BASEURL + "/pages/\\d+$": `DELETE:/pages/:id`,
	"^GET:" + config.Config.APP.BASEURL + "/pages/menus$":   `GET:/pages/menus`,

	// 接口管理
	"^POST:" + config.Config.APP.BASEURL + "/interface$":          `POST:/interface`,
	"^GET:" + config.Config.APP.BASEURL + "/interface/page/\\d+$": `GET:/interface/page/:id`,
	"^GET:" + config.Config.APP.BASEURL + "/interface/role/\\d+$": `GET:/interface/role/:id`,
	"^DELETE:" + config.Config.APP.BASEURL + "/interface/\\d+$":   `DELETE:/interface/:id`,
	"^PATCH:" + config.Config.APP.BASEURL + "/interface/\\d+$":    `PATCH:/interface/:id`,

	// 日志管理
	"^GET:" + config.Config.APP.BASEURL + "/log/commits": `GET:/log/commits`,
	"^GET:" + config.Config.APP.BASEURL + "/log/system":  `GET:/log/system`,
}
