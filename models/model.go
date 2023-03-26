package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"renWoXing/library"
)

var db *gorm.DB

type Model struct {
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
	DeleteTime int `json:"delete_time"`
}

func MySQLConnect() {
	var err error
	var conf = library.Config.DB
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	fmt.Println("dsn = ", dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
