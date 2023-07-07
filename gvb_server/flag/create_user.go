package flag

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"gvb_server/service/user_ser"
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

	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}
	err := user_ser.UserService{}.CreateUser(userName, nickName, password, role, email, "127.0.0.1")

	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功！", userName)
}
