package model

import "time"

/*
	公用件，包含id、创建时间、更新时间
*/

type Common struct {
	Id       uint64    `json:"id" gorm:"primary_key"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
