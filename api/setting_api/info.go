package setting_api

import (
	"goblog/global"
	"goblog/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingInfoView(c *gin.Context) {
	uri := SettingURI{}
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.FailWithCode(res.URIError, c)
		return
	}
	switch uri.Name {
	case "info":
		res.OkWithData(global.Config.SiteInfo, c)
	case "qq":
		res.OkWithData(global.Config.QQInfo, c)
	case "jwt":
		res.OkWithData(global.Config.JWT, c)
	default:
		res.FailWithCode(res.AugmentError, c)
	}

}
