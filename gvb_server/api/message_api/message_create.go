package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"` // 发送人id
	RevUserID  uint   `json:"rev_user_id" binding:"required"`  // 接收人id
	Content    string `json:"content" binding:"required"`      // 消息内容

}

// MessageCreateView 发布消息
func (MessageApi) MessageCreateView(c *gin.Context) {

	// 当前用户发送消息
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var senUser, recvUser models.UserModel
	err = global.DB.Take(&senUser, cr.SendUserID).Error
	if err != nil {
		res.FailWithMessage("发送人不存在", c)
		return
	}

	err = global.DB.Take(&recvUser, cr.RevUserID).Error
	if err != nil {
		res.FailWithMessage("接收人不存在", c)
		return
	}
	err = global.DB.Create(&models.MessageModel{

		SendUserID:       cr.SendUserID,
		SendUserNickName: senUser.NickName,
		SendUserAvatar:   senUser.AvatarID,
		RevUserID:        cr.RevUserID,
		RevUserNickName:  recvUser.NickName,
		RevUserAvatar:    recvUser.AvatarID,
		IsRead:           false,
		Content:          cr.Content,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("消息发送失败", c)
		return
	}
	res.OkWithMessage("消息发送成功", c)
	return
}
