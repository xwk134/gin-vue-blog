package models

import (
	"time"
)

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id" structs:"-"` //主键id
	CreatedAt time.Time `json:"created_at" structs:"-"`           //创建时间
	UpdatedAt time.Time `json:"update_at" structs:"-"`            //更新时间
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
	Role  string `form:"role"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
