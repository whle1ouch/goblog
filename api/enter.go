package api

import "goblog/api/setting_api"

type ApiGroup struct {
	SettingApi setting_api.SettingApi
}

var API = new(ApiGroup)
