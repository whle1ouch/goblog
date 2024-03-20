package router

import (
	"goblog/api"
)

func (r *RouterGroup) SettingsRouter() {
	r.GET("/setting/:name", api.API.SettingApi.SettingInfoView)
	r.PUT("/setting/:name", api.API.SettingApi.SettingInfoUpdateView)
}
