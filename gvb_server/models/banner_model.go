package models

// BannerModel banner表
type BannerModel struct {
	MODEL
	Path string `json:"path"`                //图片路径
	Hash string `json:"hash"`                //图片的hash值，用于判断重复图片
	Name string `gorm:"size:30" json:"name"` //图片名称

}
