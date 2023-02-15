package model

/*
	用户数据模型
*/

type User struct {
	Common
	Name          string `json:"name"`
	PassWord      string `json:"password"`
	FollowCount   uint64 `json:"follow_count"`
	FollowedCount uint64 `json:"followed_count"`
}
