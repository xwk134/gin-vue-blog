package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/jwts"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "zhangsan",
		NickName: "张三",
	})
	fmt.Println(token, err)

	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InpoYW5nc2FuIiwibmlja19uYW1lIjoi5byg5LiJIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNjg2NTA3NzA5LjY3MzUxOTEsImlzcyI6InRlc3QifQ.LZzNdXUkJu32aEyhTzfz9T7ezzortpQ2XG_EK9EwXKc")
	fmt.Println(claims, err)
}
