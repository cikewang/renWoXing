package models

import (
	"gorm.io/gorm"
)

type Labels struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Status   int8   `json:"status"`
	Position string `json:"position"`
}

// GetLabels 根据 admin_id 查询 admin 信息
func GetLabels() ([]Labels, error) {
	var labs []Labels
	err := db.Where("is_del = 0").Find(&labs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return labs, err
	}

	return labs, nil
}
