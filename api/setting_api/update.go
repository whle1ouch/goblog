package setting_api

import (
	"errors"
	"goblog/config"
	"goblog/core"
	"goblog/global"
	"goblog/models/res"

	"github.com/gin-gonic/gin"
)

var newMap = map[string]interface{}{
	"info": func() config.SiteInfo { return config.SiteInfo{} },
	"qq":   func() config.QQInfo { return config.QQInfo{} },
	"jwt":  func() config.JWT { return config.JWT{} },
}

func handler(uri string, c *gin.Context) (any, error) {
	cr, ok := newMap[uri]
	if !ok {
		return nil, errors.New("uri错误")
	}
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		return nil, err
	}
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (SettingApi) SettingInfoUpdateView(c *gin.Context) {
	var uri SettingURI
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.FailWithCode(res.URIError, c)
		return
	}

	cr, err := handler(uri.Name, c)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	config.Update()

	err = core.WriteConfigToYaml()
	if err != nil {
		res.FailWithCode(res.WriteSettingError, c)
		global.Log.Info("重新读取配置文件")
		core.InitConfig()
	} else {
		res.OkWithMessage("更新成功", c)
	}
}
