package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

// @title gvb_serve API文档
// @version 1.0
// @description API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	// 读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
	//初始化日志
	global.Log = core.InitLogger()
	logrus.Info("hello,world")
	//连接数据库
	global.DB = core.InitGorm()
	// 连接redis
	global.Redis = core.ConnectRedis()
	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	//初始化路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在：%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
