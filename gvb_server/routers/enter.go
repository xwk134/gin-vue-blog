package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gvb_server/global"
	"net/http"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.StaticFS("uploads", http.Dir("uploads"))
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	//系统配置api
	routerGroupApp.SettingsRouter()
	//图片管理api
	routerGroupApp.ImagesRouter()
	//广告管理api
	routerGroupApp.AdvertRouter()
	//菜单管理api
	routerGroupApp.MenuRouter()
	//用户管理api
	routerGroupApp.UserRouter()
	//标签管理api
	routerGroupApp.TagRouter()
	//消息管理api
	routerGroupApp.MessageRouter()
	//文章管理api
	routerGroupApp.ArticleRouter()
	return router
}
