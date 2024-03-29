package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/utils/desens"
	"gvb_server/utils/jwts"
)

type UserResponse struct {
	models.UserModel
	RoleID int `json:"role_id"`
}

type UserListRequest struct {
	models.PageInfo
	Role int `json:"role" form:"role"`
}

func (UserApi) UserListView(c *gin.Context) {
	// 如何判断是管理员

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var page UserListRequest
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var users []UserResponse
	list, count, _ := common.ComList(models.UserModel{Role: ctype.Role(page.Role)}, common.Option{
		PageInfo: page.PageInfo,
	})
	for _, user := range list {

		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			//管理员
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, UserResponse{
			UserModel: user,
			RoleID:    int(user.Role),
		})
	}

	res.OkWithList(users, count, c)
}
