package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
	//初始化日志
	global.Log = core.InitLogger()
	logrus.Info("hello,word")
	//连接数据库
	global.DB = core.InitGorm()
}
