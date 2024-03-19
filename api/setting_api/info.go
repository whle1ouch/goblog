package setting_api

import (
	"goblog/global"
	"goblog/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
