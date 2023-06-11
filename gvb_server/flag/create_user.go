package flag

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils"
)

func CreateUser(permissions string) {
	//用户名 昵称 密码 确认密码 邮箱
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称：")
	fmt.Scan(&nickName)
	fmt.Printf("请输入密码：")
	fmt.Scan(&password)
	fmt.Printf("请再次输入密码：")
	fmt.Scan(&rePassword)
	fmt.Printf("请输入邮箱：")
	fmt.Scan(&email)

	// 判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		global.Log.Error("用户名已存在，请重新输入")
		return
	}
	//校验两次密码
	if password != rePassword {
		global.Log.Error("两次密码不一致，请重新输入")
		return
	}
	// 对密码进行hash
	hashAndSalt := utils.HashAndSalt(password)
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}
	//头像问题
	//1、默认头像
	//2、随机选择头像
	avatar := "/uploads/file/京东.png"

	//入库
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashAndSalt,
		Email:      email,
		Role:       role,
		AvatarID:   avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功！", userName)
}
