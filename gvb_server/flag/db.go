package flag

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
)

func Makemigrations() {
	fmt.Println("开始迁移表结构")
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	// 生成表结构
	global.DB.Migrator().AutoMigrate(
		&models.AdvertModel{},
		&models.UserModel{},
		&models.CommentModel{},
		&models.ArticleModel{},
		&models.BannerModel{},
		&models.MessageModel{},
		&models.ArticleModel{},
		&models.FadeBackModel{},
		&models.TagModel{},
		&models.MenuModel{},
		&models.MenuBannerModel{},
		&models.LoginDataModel{},
	)
	if err != nil {
		global.Log.Error("生成数据表结构失败")
		return
	}
	global.Log.Error("生成数据表结构成功！")
}
