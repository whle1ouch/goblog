package setting_api

import (
	"goblog/config"
	"goblog/core"
	"goblog/global"
	"goblog/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.AugmentError, c)
		return
	}
	global.Config.SiteInfo.Update(cr)
	err = core.WriteConfigToYaml()
	if err != nil {
		res.FailWithCode(res.WriteSettingError, c)
		global.Log.Info("重新读取配置文件")
		core.InitConfig()
	} else {
		res.OkWithMessage("更新成功", c)
	}
}
