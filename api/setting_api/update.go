package setting_api

import (
	"goblog/config"
	"goblog/core"
	"goblog/global"
	"goblog/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingInfoUpdateView(c *gin.Context) {
	var uri SettingURI
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.FailWithCode(res.URIError, c)
		return
	}

	switch uri.Name {
	case "info":
		var cr config.SiteInfo
		err = c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
		}
		config.Update(&(global.Config.SiteInfo), cr)
	case "qq":
		var cr config.QQInfo
		err = c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
		}
		config.Update(&(global.Config.QQInfo), cr)
	case "jwt":
		var cr config.SiteInfo
		err = c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
		}
		config.Update(&(global.Config.SiteInfo), cr)
	default:
		res.FailWithCode(res.AugmentError, c)
		return
	}
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	err = core.WriteConfigToYaml()
	if err != nil {
		res.FailWithCode(res.WriteSettingError, c)
		global.Log.Info("重新读取配置文件")
		core.InitConfig()
	} else {
		res.OkWithMessage("更新成功", c)
	}
}
