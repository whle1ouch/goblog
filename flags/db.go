package flags

import (
	"goblog/global"
	"goblog/models"
)

func MakeMigrations() {
	global.DB.SetupJoinTable(&models.UserModel{}, "Collections", &models.UserCollectionModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "MenuBanner", &models.MenuBannerModel{})

	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},
		&models.TagModel{},
		&models.MessageModel{},
		&models.ArticleModel{},
		&models.MenuModel{},
		&models.BannerModel{},
		&models.CommentModel{},
		&models.FadeBackModel{},
		&models.LoginDataModel{},
	)
	if err != nil {
		global.Log.Error("数据库迁移失败", err)
		return
	}
	global.Log.Info("数据库迁移成功")
}
