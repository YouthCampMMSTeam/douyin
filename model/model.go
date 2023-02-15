package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "***" //dsn未确定，可用配置文件进行存储和修改，此处暂且使用这个

var (
	DB *gorm.DB
)

/*
	连接数据库，并在没有表的情况下自动创建对应表，以达到规范化
*/
func ConnectDatabase() error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := db.AutoMigrate(
		&Comment{},
		&Favorite{},
		&Message{},
		&Relation{},
		&User{},
		&Video{},
	); err != nil {
		return err
	}
	DB = db
	return nil
}
