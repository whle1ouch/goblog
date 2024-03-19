package router

import (
	"goblog/api"
)

func (r *RouterGroup) SettingsRouter() {
	r.GET("/setting/info", api.API.SettingApi.SettingInfoView)
	r.PUT("/setting/info", api.API.SettingApi.SettingInfoUpdateView)
}
