package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenRequest struct {
	MenuTitle     string      `json:"menu_title" binding:"required" msg:"请完善菜单名称"`
	MenuTitleEn   string      `json:"menu_title_en" binding:"required" msg:"请完善菜单英文名称"`
	Slogan        string      `json:"slogan"`
	Abstract      ctype.Array `json:"abstract"`
	AbstractTime  int         `json:"abstract_time"` //切换时间
	BannerTime    int         `json:"banner_time"`
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号"` //排序
	ImageSortList []ImageSort `json:"image_sort_list"`                       //具体图片的顺序
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 创建banner数据入库
	menuModel := models.MenuModel{
		MenuTitle:    cr.MenuTitle,
		MenuTitleEn:  cr.MenuTitleEn,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败", c)
		return
	}

	if len(cr.ImageSortList) == 0 {
		res.OkWithMessage("菜单添加成功", c)
		return
	}

	var menuBannerList []models.MenuBannerModel
	for _, sort := range cr.ImageSortList {
		// 判断image_id是否真正有这张图片
		menuBannerList = append(menuBannerList, models.MenuBannerModel{

			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}

	//给第三张表入库

	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}
	res.OkWithMessage("菜单添加成功", c)
}