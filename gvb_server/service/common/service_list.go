package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Debug().Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" //默认按时间往前排序
	}

	if option.Key != "" {
		count = DB.Debug().Select("id").Where("user_name like ? or nick_name like ?", option.Key+"%", option.Key+"%").Find(&list).RowsAffected

	}
	if option.Key == "" {
		count = DB.Select("id").Find(&list).RowsAffected
	}

	if option.Role != "" {
		count = DB.Debug().Select("id").Where("role = ? ", option.Role).Find(&list).RowsAffected

	}

	if option.Key != "" && option.Role != "" {
		count = DB.Debug().Select("id").Where("user_name like ? or nick_name like ? and role = ?", option.Key+"%", option.Key+"%", option.Role).Find(&list).RowsAffected

	}

	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}

	if option.Key != "" {
		err = DB.Debug().Limit(option.Limit).Offset(offset).Order(option.Sort).Where("user_name like ? or nick_name like ?", option.Key+"%", option.Key+"%").Find(&list).Error

	}
	if option.Key == "" {
		err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	}
	if option.Role != "" {
		err = DB.Debug().Limit(option.Limit).Offset(offset).Order(option.Sort).Where("role = ? ", option.Role).Find(&list).Error

	}

	if option.Key != "" && option.Role != "" {
		err = DB.Debug().Limit(option.Limit).Offset(offset).Order(option.Sort).Where("user_name like ? or nick_name like ? and role = ?", option.Key+"%", option.Key+"%", option.Role).Find(&list).Error

	}

	return list, count, err
}
