package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Action struct {
	Id         int `json:"id"`
	UpdateTime int `json:"update_time"`
}

// GetUpDateTime 获取更新时间
func GetUpDateTime() (Action, error) {
	var act Action
	err := db.Where("id = 1").Find(&act).Error
	fmt.Println(err)
	if err != nil && err != gorm.ErrRecordNotFound {
		return act, err
	}

	return act, nil
}
