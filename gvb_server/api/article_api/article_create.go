package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type ArticleRequest struct {
	Title			string		`json:"title" binding:"required" msg:"请输入文章标题" structs:"title"` //文章的标题
	Abstract		string		`json:"abstract" binding:"required" msg:"请输入文章简介" structs:"abstract"` //文章简介
	Content			string		`json:"content" binding:"required" msg:"请输入文章内容" structs:"content"` //文章内容
	LookCount		int			`json:"look_count" binding:"required" msg:"请输入文章浏览量" structs:"look_count"` //浏览量
	CommentCount	int			`json:"comment_count" binding:"required" msg:"请输入文章评论量" structs:"comment_count"` //评论量
	DiggCount		int			`json:"digg_count" binding:"required" msg:"请输入文章点赞量" structs:"digg_count"` //点赞量
	CollectsCount	int 		`json:"collects_count" binding:"required" msg:"请输入文章收藏量" structs:"collects_count"` //收藏量
	
	UserID			uint		`json:"user_id" binding:"required" msg:"请输入用户id" structs:"user_id"` //用户id
	Category		string		`json:"category" binding:"required" msg:"请输入文章分类" structs:"category"` //文章分类
	Source			string		`json:"source" binding:"required" msg:"请输入文章来源" structs:"source"` //文章来源
	Link			string		`json:"link" structs:"link"` //原文链接
			
	BannerID		uint		`json:"banner_id" binding:"required" msg:"请输入封面id" structs:"banner_id"` //封面id
	NickName		string		`json:"nick_name" binding:"required" msg:"请输入发布文章的用户昵称" structs:"nick_name"` //发布文章的用户昵称
	
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 重复判断
	var article models.ArticleModel
	err = global.DB.Take(&article, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该文章已存在", c)
		return
	}
	err = global.DB.Debug().Create(&models.ArticleModel{
		Title: cr.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加文章失败", c)
	}
	res.OkWithMessage("添加文章成功", c)
}
