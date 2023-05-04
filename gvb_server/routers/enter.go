package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	//系统配置api
	routerGroupApp.SettingsRouter()
	//图片管理api
	routerGroupApp.ImagesRouter()
	//广告管理api
	routerGroupApp.AdvertRouter()
	return router
}
